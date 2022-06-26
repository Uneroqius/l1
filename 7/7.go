package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	nums := []float64{2, 4, 6, 8, 10}
	sqrtNums := map[float64]float64{}

	for _, num := range nums {
		wg.Add(1)

		//
		go func(num float64, m map[float64]float64, mutex *sync.Mutex) {
			sqrt := math.Sqrt(float64(num))
			mutex.Lock()
			m[num] = sqrt
			mutex.Unlock()
			wg.Done()
		}(num, sqrtNums, &mutex)
	}

	wg.Wait()
	for _, num := range nums {
		fmt.Print(sqrtNums[num], " ")
	}
	fmt.Println()
}
