package main

import "fmt"

func main() {
	a := 14

	fmt.Println("a - ", a)
	fmt.Println("a's memory adres", &a)
	fmt.Printf("%d \n", &a)
}
