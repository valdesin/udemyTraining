package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	incrementA := wrapper()
	incrementB := wrapper()

	fmt.Println(incrementA())
	fmt.Println(incrementA())
	fmt.Println(incrementB())
	fmt.Println(incrementB())
	fmt.Printf("%T \n", incrementB)
}
