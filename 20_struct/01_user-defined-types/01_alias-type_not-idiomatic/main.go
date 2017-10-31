package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 23
	s := myAge + 23
	fmt.Printf("%T %v \n", myAge, myAge)
	fmt.Printf("%T \n", s)
}
