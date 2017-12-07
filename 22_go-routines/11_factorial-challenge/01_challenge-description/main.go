package main

import "fmt"

func main() {
	f := fuctorial(4)
	fmt.Println(f)
}

func fuctorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
