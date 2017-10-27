package main

import "fmt"

func main() {
	v, even := half(4)
	fmt.Println(v, " ", even)
}

func half(v int) (int, bool) {
	return v / 2, v%2 == 0
}
