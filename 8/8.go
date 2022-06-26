package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		num         int64
		bitPosition int64
		val         int64
	)

	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			main()
		}
	}()

	fmt.Print("Input the number: ")
	fmt.Scan(&num)
	fmt.Print("Input the bit position (1-64): ")
	fmt.Scan(&bitPosition)
	if bitPosition > 64 {
		panic("is too big")
	} else if bitPosition < 1 {
		panic("is too small")
	}
	fmt.Print("Input the bit value: ")
	fmt.Scan(&val)
	if val != 0 && val != 1 {
		panic("1 or 0")
	}

	// выводим состояния битов исходного числа
	fmt.Printf("%#.64b\n", uint64(num))

	// создаем маску числа, состояющую из 63 нулей и
	// одной единицы на bitPosition-ой позиции
	sieve := int64(1) << (bitPosition - 1)

	// обнуляем бит на позиции bitPosition
	num &^= sieve

	// создаем маску числа, состояющую из 63 нулей и
	// одного(-ой) нуля(единицы), в зависимости от значения,
	// к которому нужно приравнять бит на позиции bitPosition
	sieve = val << (bitPosition - 1)

	// записываем бит на позиции bitPosition в нужное значение
	num |= sieve

	// Выводим состояния битов измененного числа
	fmt.Printf("%#.64b\n", uint64(num))
	fmt.Println(num)
}
