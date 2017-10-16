package main

import "fmt"

func main() {
	fmt.Println(greet("Vlad ", "Fill"))
}

func greet(name, lastName string) (s string) {
	s = fmt.Sprint(name, lastName)
	return
}
