package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Square struct {
	sideLength float64
}

type Triangle struct {
	height float64
	base   float64
}

func (square Square) getArea() float64 {
	return square.sideLength * square.sideLength
}

func (triangle Triangle) getArea() float64 {
	return .5 * triangle.base * triangle.height
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {
	square := Square{sideLength: 10}
	triangle := Triangle{base: 10, height: 10}

	printArea(square)
	printArea(triangle)
}
