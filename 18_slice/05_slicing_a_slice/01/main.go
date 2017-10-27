package main

import "fmt"

func main() {

	var mySlice []int
	fmt.Println(mySlice)

	sSlice := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(sSlice)
	fmt.Println(sSlice[1:5])
	fmt.Println(sSlice[5])
	fmt.Println("sSlice"[1])
}
