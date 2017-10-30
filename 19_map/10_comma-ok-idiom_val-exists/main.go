package main

import "fmt"

func main() {
	greeting := map[int]string{
		0: "Hello",
		1: "Pryvit",
		2: "Hola",
		3: "Bonjorno",
	}
	delete(greeting, 2)

	if v, exist := greeting[2]; exist {
		fmt.Println("This val exist: ")
		fmt.Println("Val:", v)
		fmt.Println("Exist:", exist)
	} else {
		fmt.Println("This val did't exist: ")
		fmt.Println("Val:", v)
		fmt.Println("Exist:", exist)
	}

	fmt.Println(greeting)
}
