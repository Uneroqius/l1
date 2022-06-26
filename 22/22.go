package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b big.Float
	fmt.Print("Input a: ")
	fmt.Scan(&a)
	fmt.Print("Input b: ")
	fmt.Scan(&b)

	fmt.Println("a * b = ", new(big.Float).Mul(&a, &b))
	fmt.Println("a / b = ", new(big.Float).Quo(&a, &b))
	fmt.Println("a + b = ", new(big.Float).Add(&a, &b))
	fmt.Println("a - b = ", new(big.Float).Mul(&a, &b))
}
