package main

import "fmt"

func main() {
	switch "Vlad" {
	case "Vlad":
		fmt.Println("Wasup Vlad")
	case "Daniel":
		fmt.Println("Wasup Daniel")
	case "Vita":
		fmt.Println("Vita")
	default:
		fmt.Println("Not found this name.")
	}
}
