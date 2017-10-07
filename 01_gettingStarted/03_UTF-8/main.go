package main

import "fmt"

func main() {

	for i := 60; i < 122; i++ {
		fmt.Printf("%d - %b - %x - %q \n", i, i, i, i)
	}
}
