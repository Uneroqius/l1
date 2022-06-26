package main

import (
	"fmt"
	"unicode"
)

func reverseWord(word []rune) {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
}

func skipSpaces(s []rune, i *int) {
	for *i < len(s) && unicode.IsSpace(s[*i]) {
		*i += 1
	}
}

func skipNonSpaces(s []rune, i *int) {
	for *i < len(s) && !unicode.IsSpace(s[*i]) {
		*i += 1
	}
}

func main() {
	s1 := "snow dog sun"

	rS1 := []rune(s1)
	var beginOfWordIndx int

	for i := 0; i < len(rS1); {
		skipSpaces(rS1, &i)
		beginOfWordIndx = i
		skipNonSpaces(rS1, &i)
		reverseWord(rS1[beginOfWordIndx:i])
	}

	s1 = string(rS1)

	fmt.Println(s1)
}
