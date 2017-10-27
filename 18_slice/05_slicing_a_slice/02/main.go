package main

import "fmt"

func main() {
	sSlice := []string{
		"Good morning!",
		"Bonjour!",
		"Buenos dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	fmt.Print("[1:5]")
	fmt.Println(sSlice[1:5])
	fmt.Print("[0:2]")
	fmt.Println(sSlice[0:2])
	fmt.Print(":4")
	fmt.Println(sSlice[:4])
	fmt.Print(":")
	fmt.Println(sSlice[:])
}
