package main

import "fmt"

func main() {
	fmt.Println(greet("Vlad", "Filonenko"))
}

func greet(name, lastName string) string {
	return fmt.Sprint(name, lastName)
}
