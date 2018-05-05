package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{sideLength: 12.0}
	tri := triangle{base: 12.0, height: 10.1}

	printArea(sq)
	printArea(tri)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (tr triangle) getArea() float64 {
	return tr.base * tr.height * 0.5
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
