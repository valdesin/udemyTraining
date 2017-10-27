package main

import "fmt"

func main() {
	sSlice := make([]string, 3, 5)

	sSlice[0] = "Hello"
	sSlice[1] = "Buenos"
	sSlice[2] = "Pryvit"
	sSlice = append(sSlice, "Alohaa")

	fmt.Println(sSlice[3])
}
