package main

import "fmt"

func main() {
	var greeting = make(map[string]string)

	greeting["Tim"] = "Bonjorno"
	greeting["Vlad"] = "Pryvit"

	fmt.Println(greeting)
}
