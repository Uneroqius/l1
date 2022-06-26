package main

import (
	"fmt"
)

func main() {
	a, b := 1, 4
	a, b = b, a

	fmt.Println(a, b)
}
