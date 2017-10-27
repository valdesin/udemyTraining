package main

import "fmt"

func main() {
	var x [52]int

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[14])
	x[14] = 14
	fmt.Println(x[14])
}
