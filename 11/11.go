package main

import (
	"fmt"
)

// каждое значение из слайсов добавляется в мапу m
// одиннаковых ключей в мапе быть не может
// записываем все ключи в set3 - это будет пересечение двух множеств
func main() {
	set1 := []int{1, 11, 3, 5, 9, 10, 44, -10, 55, 32}
	set2 := []int{-100, -200, -10, 4, 5, 55}
	m := map[int]struct{}{}

	set3 := []int{}
	for _, elem := range set1 {
		m[elem] = struct{}{}
	}
	for _, elem := range set2 {
		if _, isExist := m[elem]; isExist {
			set3 = append(set3, elem)
		}
	}
	fmt.Println(set3)
}
