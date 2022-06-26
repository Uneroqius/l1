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
	sum := 0

	for _, num := range nums {
		wg.Add(1)
		go func(num int, m *map[int]int, mutex *sync.RWMutex) {
			square := num * num

			// Блокируем операции изменения мапы, чтобы избежать гонку данных
			mutex.Lock()
			(*m)[num] = square
			// Разблокировка
			mutex.Unlock()
			wg.Done()
		}(num, &sqrtNums, &mutex)
	}

	wg.Wait()

	for _, num := range nums {
		sum += sqrtNums[num]
	}
	fmt.Println(sum)
}
