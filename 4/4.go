package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

var workerCount int

func init() {
	// Инициализация флага для реализации возможности указывать количество воркеров
	flag.IntVar(&workerCount, "worker-count", 6, "количество воркеров")
}

type worker func(int, chan interface{}, chan interface{})

var workers []worker

func main() {
	flag.Parse()

	ch := make(chan interface{})
	finCh := make(chan struct{})
	wg := sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)

		go func(ID int, ch chan interface{}, finCh chan struct{}) {
			defer func() {
				// перед завершением работы горутины
				// выводим сообщение
				fmt.Printf("Worker %d is finished\n", ID)
				// и уменьшаем счетчик горутин для ожидания
				wg.Done()
			}()

			for {
				// смотрим, из какого канал пришло сообщение
				select {
				case elem := <-ch:
					// Сообщение пришло из канала с данными
					// Выводим на экран
					fmt.Printf("Worker %d: %v\n", ID, elem)
					time.Sleep(time.Second)
				case <-finCh:
					return
				}
			}
		}(i+1, ch, finCh)
	}

	inputCh := make(chan interface{})
	// горутина, занимающаяся чтением данных с консоли от отправкой этих данных в канал inputCH
	go func(ch chan interface{}) {
		var s string
		for {
			fmt.Scan(&s)
			ch <- s
		}
	}(inputCh)

	// после нажатия клавиш crtl+c сигнал посылается в канал stopSig
	stopSig := make(chan os.Signal)
	signal.Notify(stopSig, os.Interrupt)

	defer func() {
		// перед завершением рабобы главной горутины
		// отправляем остальным горутинам сигнал с завершением работы
		for i := 0; i < workerCount; i++ {
			finCh <- struct{}{}
		}
		wg.Wait()
		fmt.Println("Main worker is finished")
	}()

	for {
		select {
		case <-stopSig:
			return
		case s := <-inputCh:
			ch <- s
		}
	}
}
