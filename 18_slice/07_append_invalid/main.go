package main

import "fmt"

func main() {
	sSlice := make([]string, 3, 5)

	sSlice[0] = "Hello"
	sSlice[1] = "Bueno"
	sSlice[2] = "Pryvit"
	sSlice[3] = "Alohaa"

	fmt.Println(sSlice[2])
}
