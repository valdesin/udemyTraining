package main

import "fmt"

func main() {
	a := 14

	fmt.Println(a)
	fmt.Println(&a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b)
}
