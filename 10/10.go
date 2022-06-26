package main

import (
	"fmt"
	"sort"
)

// находим ближайший наименьший по модулю десяток
func findMinDozen(num float64) (step float64) {
	if num <= -10 {
		step = -10
		for !(step > num && step-10 < num) {
			step -= 10
		}
	} else if num >= 10 {
		step = 10
		for !(step < num && step+10 > num) {
			step += 10
		}
	} else {
		step = 0
	}

	return
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 0.10, 15.5, 24.5, -21.0, 32.5}
	res := map[float64]*[]float64{}
	var step float64

	sort.Float64s(temps)
	for _, temp := range temps {
		// для каждого числа находит ближайший минимальный по модулю десяток
		step = findMinDozen(temp)
		sl, isExist := res[step]
		// если ключа в мапе по этому десятку еще нет
		if !isExist {
			// инициализируем слайс
			res[step] = &[]float64{}
			sl = res[step]
		}
		// добавляем температуру в нужный слайс
		*sl = append(*sl, temp)
	}

	isFirstSl := true
	for dz, sl := range res {
		if !isFirstSl {
			fmt.Print(", ")
		} else {
			isFirstSl = false
		}
		fmt.Printf("%.0f:{", dz)
		for i, temp := range *sl {
			fmt.Print(temp)
			if i != len(*sl)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("}")
	}
	fmt.Println()
}
