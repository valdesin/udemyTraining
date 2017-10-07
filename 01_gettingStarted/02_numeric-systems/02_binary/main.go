package main

import "fmt"

func main() {
	fmt.Printf("%d - %b \n", 42, 42)

	fmt.Printf("%s - %x \n", "A", "A")

	for i := 0; i < 100; i++ {
		fmt.Printf("%d - %b \n", i, i)
	}
}
