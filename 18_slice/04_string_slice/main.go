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

	for i, v := range sSlice {
		fmt.Println(i, " - ", v)
	}

	for j := 0; j < len(sSlice); j++ {
		fmt.Println(sSlice[j])
	}
}
