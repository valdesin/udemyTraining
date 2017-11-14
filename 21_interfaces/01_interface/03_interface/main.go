package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
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

func main() {
	s := square{10}
	c := circle{7}

	fmt.Printf("%T \n", s)
	fmt.Printf("%T \n", c)

	info(s)
	info(c)
}
