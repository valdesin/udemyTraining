package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
	changeMe(&x)
	fmt.Println(&x)
	fmt.Println(x)
}

func changeMe(n *int) {
	fmt.Println(n)
	fmt.Println(*n)
	*n = 14
	fmt.Println(n)
	fmt.Println(*n)
}
