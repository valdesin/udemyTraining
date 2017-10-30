package main

import "fmt"

func main() {
	greeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	for k, v := range greeting {
		fmt.Println(k, " - ", v)
	}
}
