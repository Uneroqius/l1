package main

import (
	"fmt"
)

// дает ответ, параметр какого типа был передан
// в функцию
func Type(val interface{}) (res string) {
	switch val.(type) {
	case int:
		res = "int"
	case string:
		res = "string"
	case bool:
		res = "bool"
	default:
		res = fmt.Sprintf("%T", val)
		if res[:4] == "chan" {
			res = "channel"
		}
	}
	return
}

func main() {
	var v interface{} = 32
	fmt.Println(Type(v))

	v = "32"
	fmt.Println(Type(v))

	v = false
	fmt.Println(Type(v))

	v = make(chan interface{})
	fmt.Println(Type(v))
}
