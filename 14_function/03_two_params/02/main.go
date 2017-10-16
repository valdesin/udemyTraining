package main

import "fmt"

func greet(name, lastname string) {
	fmt.Println(name, lastname)
}

func main() {
	greet("Vlad", "Fill")
}
