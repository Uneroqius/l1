package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

var secondsCount int

func init() {
	// инициализация флага для установки времени выполнения программы
	flag.IntVar(&secondsCount, "seconds-count", 5, "длительность таймера (в секундах)")
}

func main() {
	flag.Parse()

	ch := make(chan interface{})
	wg := sync.WaitGroup{}
	finCh := make(chan struct{})
	// таймер пошлет через secondsCount секунд с начала работы программы сигнал о завершении работы
	timer := time.NewTimer(time.Second * time.Duration(secondsCount))

	wg.Add(1)
	go func(ch <-chan interface{}, finCh chan struct{}) {
		defer func() {
			fmt.Println("Goroutine closing..")
			wg.Done()
		}()

		for {
			select {
			case val := <-ch:
				fmt.Println("Goroutine: ", val)
			case <-finCh:
				return
			}
		}
	}(ch, finCh)

	// считывает данные с консоли
	go func(ch chan<- interface{}) {
		var val string
		for {
			fmt.Scan(&val)
			ch <- val
		}
	}(ch)

	defer func() {
		// посылаем сигнал о завершении побочной горутине
		finCh <- struct{}{}
		wg.Wait()
		close(finCh)
		fmt.Println("Programm closing..")
	}()

	<-timer.C
	return
}
