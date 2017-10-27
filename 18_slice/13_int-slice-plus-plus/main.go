package main

import "fmt"

func main() {
	slice := make([]int, 1)
	slice[0] = 7
	fmt.Println(slice[0])
	slice[0]++
	fmt.Println(slice[0])
}
