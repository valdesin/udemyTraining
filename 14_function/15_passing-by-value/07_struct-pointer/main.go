package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c := customer{"Vlad", 23}
	fmt.Println(&c)
	fmt.Println(c.name)
	changeMe(&c)
	fmt.Println(c.name)
}

func changeMe(c *customer) {
	fmt.Println(*c)
	fmt.Println(c.name)
	c.name = "Fill"
	fmt.Println(c.name)
}
