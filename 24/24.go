package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// конструктор
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p1 *Point) Distance(p2 *Point) (distance float64) {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	p1, p2 := NewPoint(81, 44.6), NewPoint(-20, -14.3)
	fmt.Println(p1.Distance(p2))
}
