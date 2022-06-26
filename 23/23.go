package main

import (
	"fmt"
)

func delElemFromSlice(sl *[]int, i int) {
	*sl = append((*sl)[:i], (*sl)[i+1:]...)
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	delElemFromSlice(&sl, 8)
	fmt.Println(sl)
}
