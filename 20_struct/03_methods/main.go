package main

import "fmt"

type person struct {
	first  string
	second string
	age    int
}

func (p person) fullName() string {
	return p.first + " - " + p.second
}

func main() {
	p1 := person{
		first:  "Vlad",
		second: "Fill",
		age:    23,
	}

	p2 := person{
		first:  "Todd",
		second: "Wilyam",
		age:    44,
	}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
