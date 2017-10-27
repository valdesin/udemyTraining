package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "Hello"
	greeting[1] = "Buenos"
	greeting[2] = "Pryvit"
	greeting = append(greeting, "Alohaa")
	greeting = append(greeting, "Ni hao")
	greeting = append(greeting, "konnichiwa")
	greeting = append(greeting, "Hella")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
