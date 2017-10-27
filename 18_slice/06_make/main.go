package main

import "fmt"

func main() {
	iSlice := make([]int, 3)

	iSlice[0] = 10
	iSlice[1] = 14
	iSlice[2] = 20

	fmt.Println(iSlice[0])
	fmt.Println(iSlice[1])
	fmt.Println(iSlice[2])

	sSlice := make([]string, 3, 5)

	sSlice[0] = "Hello"
	sSlice[1] = "Buenos dias"
	sSlice[2] = "Pryvet"

	fmt.Println(sSlice[0])
	fmt.Println(sSlice[1])
	fmt.Println(sSlice[2])
}
