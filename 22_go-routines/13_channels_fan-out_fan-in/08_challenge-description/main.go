package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	xc := fanOUT(in, 10)

	for v := range marge(xc...) {
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

func fanOUT(c <-chan int, n int) []<-chan int {
	var ch []<-chan int
	for i := 0; i < n; i++ {
		ch = append(ch, factorial(c))
	}
	return ch
}

func factorial(c <-chan int) <-chan int {
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
