package main

import (
	"fmt"
	"time"
)

// выполняет пустой цикл, пока
// время не достигнет нужного значения
func mySleep(d time.Duration) {
	Till := time.Now().Add(d)
	for time.Now().Before(Till) {
	}
}

func main() {
	mySleep(time.Second * 6)
	fmt.Println("Program completed")
}
