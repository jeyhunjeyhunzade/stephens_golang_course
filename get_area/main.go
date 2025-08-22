package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base float64
}

type shape interface{
	getArea() float64
}

func main() {
	square := square{sideLength: 4.0}
	triangle := triangle{height: 4, base: 3}

	printArea(square)
	printArea(triangle)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}