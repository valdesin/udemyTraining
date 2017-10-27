package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 67, 7, 8, 9, 0}

	for i, v := range sl {
		fmt.Println(i, " - ", v)
	}
}
