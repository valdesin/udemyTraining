package main

import "fmt"

func main() {
	switch "Vlad" {
	case "Volody", "Vlad":
		fmt.Println("Wassup Valy or Vlad")
	case "Daniel", "Dan":
		fmt.Println("Hello Daniel or Dan")
	case "Anton", "Andrey":
		fmt.Println("Hello Anton or Andrey")
	default:
		fmt.Println("Name NOT FOUND!!!")
	}
}
