package main

import "fmt"

func zero(z *int) {
	fmt.Println(z)
	*z = 0
}

func main() {
	x := 10
	zero(&x)
	fmt.Println(x)
}
