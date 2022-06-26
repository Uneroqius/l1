package main

import (
	"fmt"
)

func main() {
	var s string
	var r []rune

	fmt.Print("Input the word:\t")
	fmt.Scan(&s)

	r = []rune(s)

	// двигаем переменные с конца и начала строки в центр, пока i < j, меняем местами соответствующие элементы
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	s = string(r)
	fmt.Println("Result:\t", s)
}
