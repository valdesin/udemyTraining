package main

import "fmt"

func main() {
	c := fuctorial(5)

	for v := range c {
		fmt.Println(v)
	}
}

func fuctorial(n int) chan int {
	out := make(chan int)
	total := 1
	go func() {
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
