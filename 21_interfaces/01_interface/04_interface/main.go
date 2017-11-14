package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func totalArea(shapes ...shape) {
	var sum float64
	for _, v := range shapes {
		sum += v.area()
	}
	fmt.Println(sum)
}

func main() {
	s := square{10}
	c := circle{6}
	fmt.Printf("%T \n", s)
	fmt.Printf("%T \n", c)

	info(s)
	info(c)

	totalArea(s, c)
}
