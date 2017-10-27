package main

import "fmt"

func main() {
	half := func(v int) (float64, bool) {
		return float64(v) / 2, v%2 == 0
	}

	v, even := half(5)
	fmt.Println(v, even)

}
