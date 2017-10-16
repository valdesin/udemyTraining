package main

import "fmt"

func greet(name string, lastname string) {
	fmt.Println(name, lastname)
}

func main() {
	greet("Vlad", "Fill")
}
