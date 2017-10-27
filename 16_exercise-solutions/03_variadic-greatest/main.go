package main

import "fmt"

func max(num ...int) int {
	var m int
	for _, n := range num {
		if m < n {
			m = n
		}
	}
	return m
}
func main() {
	numbers := []int{1, 45, 76, 3, 88, 99, 100, 4, 5}
	m := max(numbers...)
	fmt.Println("The greatest number of list is: ", m)
}
