package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	v := []int{}

	for i := 0; i < 20; i++ {
		v = append(v, rand.Int()%100)
	}
	fmt.Println("Before:\t", v)
	// Сортирует значения типа int в порядке возрастания
	sort.Ints(v)
	fmt.Println("After:\t", v)

}
