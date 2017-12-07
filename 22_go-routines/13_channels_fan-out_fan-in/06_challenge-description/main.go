package main

import "fmt"

func main() {
	in := gen()

	f := fuctorial(in)
	for v := range f {
		fmt.Println(v)
	}
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

func fuctorial(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range c {
			out <- fact(v)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
