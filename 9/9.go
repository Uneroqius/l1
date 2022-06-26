package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// получаем из канала chIn значение int,
// вычисляем x*2, пишем его в chOut
func f1(chIn <-chan int, chOut chan<- int) {
	for x := range chIn {
		square := x * 2
		chOut <- square
	}
}

// из канала chIn получаем число int
// выводим его в stdout
func f2(chIn <-chan int) {
	for val := range chIn {
		fmt.Println(val)
		wg.Done()
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go f1(ch1, ch2)
	go f2(ch2)

	wg.Add(10)
	for _, num := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		ch1 <- num
	}

	wg.Wait()
	close(ch1)
	close(ch2)
}
