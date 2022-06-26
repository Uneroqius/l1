package main

import (
	"fmt"
)

// каждое значение из слайса sl добавляется в мапу m
// одиннаковых ключей в мапе быть не может
// пишем все ключи в res - это собственное множество
func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	res := []string{}
	m := map[string]struct{}{}

	for _, elem := range sl {
		m[elem] = struct{}{}
	}

	for elem := range m {
		res = append(res, elem)
	}

	fmt.Println(res)
}
