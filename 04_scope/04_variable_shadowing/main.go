package main

import "fmt"

func max(x int) int {
	return x + 42
}

func main() {
	max := max(7) // Its shadowing variable
	//fmt.Println(max())
	fmt.Println(max)
}
