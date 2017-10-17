package main

import "fmt"

func filte(numbers []int, callback func(n int) bool) []int {
	arrFilt := []int{}
	for _, n := range numbers {
		if callback(n) {
			arrFilt = append(arrFilt, n)
		}
	}
	return arrFilt
}

func main() {
	arrFilts := filte([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
		return n > 4
	})

	fmt.Println(arrFilts)
}
