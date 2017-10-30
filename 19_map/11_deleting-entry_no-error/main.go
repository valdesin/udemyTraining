package main

import "fmt"

func main() {
	greeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(greeting)
	delete(greeting, 7)
	fmt.Println(greeting)
}
