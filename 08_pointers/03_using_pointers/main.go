package main

import "fmt"

func main() {
	a := 14
	fmt.Println(a)
	fmt.Println(&a)

	var b = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 50
	fmt.Println(a)
}
