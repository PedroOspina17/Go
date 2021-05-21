package main

import "fmt"

func main1() {
	t := triangle{
		base:   2,
		height: 2,
	}
	printArea(t)

	s := square{
		sideLength: 2,
	}
	printArea(s)
}

func printArea(s shape) {
	fmt.Println("The area is: ", s.getArea())
}

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

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (t square) getArea() float64 {
	return t.sideLength * t.sideLength
}
