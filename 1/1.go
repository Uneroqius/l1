package main

import (
	"fmt"
)

type Action struct{}

func (a *Action) Run() {
	fmt.Println("Running")
}

func (a *Action) Jump() {
	fmt.Println("Jumping")
}

type Human struct {
	Action
	Age        int
	Name       string
	Profession string
}

func main() {
	h := Human{}
	h.Run()
	h.Jump()
}
