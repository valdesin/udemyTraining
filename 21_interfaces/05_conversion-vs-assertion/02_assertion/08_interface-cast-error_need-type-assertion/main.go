package main

import "fmt"

func main() {
	var x = 34.456
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", int(x))

	var val interface{} = 78
	fmt.Printf("%T\n", val)
	//fmt.Printf("%T\n", int(val)) //error cannot convert val (type interface {}) to type int: need type assertion
	fmt.Printf("%T\n", val.(int))
}
