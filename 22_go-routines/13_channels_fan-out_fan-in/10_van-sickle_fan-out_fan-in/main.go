package main

import (
	"fmt"
)

func main() {
	in := gen()
	out := make(chan int)

	fanOUT(in, 10, out)

	go func() {
		for v := range out {
			fmt.Println(v)
		}
	}()

	var s string
	fmt.Scanln(&s)
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fanOUT(in <-chan int, n int, out chan<- int) {
	for i := 0; i < n; i++ {
		factorial(in, out)
	}
}

func factorial(in <-chan int, out chan<- int) {
	go func() {
		for v := range in {
			out <- fact(v)
		}
	}()
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
