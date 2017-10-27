package main

import "fmt"

func main() {
	v, even := half(5)
	fmt.Println(v, " ", even)
}

func half(v int) (float64, bool) {
	return float64(v) / 2, v%2 == 0
}
