package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	letters := map[rune]struct{}{}

	fmt.Scan(&s)
	s = strings.ToLower(s)

	for _, letter := range []rune(s) {
		// если буква уже есть в мапе, то значит,
		// что она повторяется
		if _, isExist := letters[letter]; isExist {
			fmt.Println("false")
			return
		}
		letters[letter] = struct{}{}
	}

	fmt.Println("true")
}
