package main

import "fmt"

func main() {
	myName := "Vlad"

	switch {
	case len(myName) == 2:
		fmt.Println("Length = 2")
	case myName == "Volodymer":
		fmt.Println("Wassup Volodymer")
	case myName == "Vlad":
		fmt.Println("Wassup Vlad")
	case myName == "Anton":
		fmt.Println("Wassup Anton")
	default:
		fmt.Println("Name NOT FOUND")
	}
}
