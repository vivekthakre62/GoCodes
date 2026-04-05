package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

// Circle methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Length: 4, Width: 3}

	var s Shape

	s = c
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())

	s = r
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())
}
