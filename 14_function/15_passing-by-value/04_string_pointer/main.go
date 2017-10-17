package main

import "fmt"

func main() {
	name := "Vlad"
	fmt.Println(&name)
	fmt.Println(name)
	changeMe(&name)
	fmt.Println(&name)
	fmt.Println(name)
}

func changeMe(n *string) {
	fmt.Println(n)
	fmt.Println(*n)
	*n = "Fill"
	fmt.Println(n)
	fmt.Println(*n)
}
