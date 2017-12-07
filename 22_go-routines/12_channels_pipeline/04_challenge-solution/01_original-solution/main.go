package main

import "fmt"

func main() {
	in := gen()
	out := factorial(in)
	for v := range out {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(c <-chan int) <-chan int {
	out := make(chan int)

	fact := func(n int) int {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		return total
	}

	go func() {
		for v := range c {
			out <- fact(v)
		}
		close(out)
	}()
	return out
}
