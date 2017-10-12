package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p", &z)
	fmt.Println(&z)
	z = 0
}

func main() {
	x := 10
	fmt.Printf("%p", &x)
	fmt.Println(&x)
	zero(x)
	fmt.Println(x)
}
