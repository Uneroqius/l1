package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mutex := sync.RWMutex{}
	nums := []int{2, 4, 6, 8, 10}
	sqrtNums := map[int]int{}

	for _, num := range nums {
		wg.Add(1)
		go func(num int, m *map[int]int, mutex *sync.RWMutex) {
			sqrt := num * num
			// Блокируем операции изменения мапы, чтобы избежать гонку данных
			mutex.Lock()
			(*m)[num] = sqrt
			// Разблокировка
			mutex.Unlock()
			wg.Done()
		}(num, &sqrtNums, &mutex)
	}

	wg.Wait()

	for _, num := range nums {
		fmt.Print(sqrtNums[num], " ")
	}
	fmt.Println()
}
