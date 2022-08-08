package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	s := square{sideLength: 5.5}
	t := triangle{base: 5.4, height: 6.5}

	printArea(s)
	printArea(t)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	a := s.getArea()
	fmt.Println("Area is ", a)
}
