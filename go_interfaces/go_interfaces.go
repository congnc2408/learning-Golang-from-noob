package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (r circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

func (r circle) perim() float64 {
	return 2 * math.Pi * r.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius: ", c.radius)
	}
	fmt.Println("có cức")
}

func main() {
	r := rect{width: 10, height: 50}
	c := circle{radius: 24}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
