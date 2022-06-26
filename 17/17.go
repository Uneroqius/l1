package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	sl := []int{}

	for i := 0; i < 20; i++ {
		sl = append(sl, rand.Int()%100)
	}

	// сортируем
	sort.Ints(sl)
	numToFind := sl[10]

	// бинарный поиск
	index := sort.SearchInts(sl, numToFind)
	fmt.Println("After:\t", index)

}
