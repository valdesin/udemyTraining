package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	for v := range marge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
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

func factorial(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range out {
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

func marge(ch ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for v := range c {
			out <- v
		}
		wg.Done()
	}

	wg.Add(len(ch))

	for _, v := range ch {
		go output(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
