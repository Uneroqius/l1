package main

import (
	"fmt"
	"time"
)

func main() {
	// Естественное завершение работы функции в горутине
	fmt.Println("Естественное завершение работы функции в горутине")
	go func() {
		fmt.Println("Goroutine is finished...")
	}()
	time.Sleep(time.Second)
	fmt.Println()

	// Завершение работы горутины с помощью передачи сообщения через канал
	fmt.Println("Завершение работы горутины с помощью передачи сообщения через канал")
	ch := make(chan interface{})
	go func(ch <-chan interface{}) {
		<-ch
		fmt.Println("Goroutine is finished...")
	}(ch)
	ch <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println()

	// Завершение работы горутины по таймеру
	fmt.Println("Завершение работы горутины по таймеру")
	timer := time.NewTimer(time.Second)
	go func(timer *time.Timer) {
		<-timer.C
		fmt.Println("Goroutine is finished...")
	}(timer)
	time.Sleep(time.Second * 2)
	fmt.Println()
}
