package main

import "fmt"

func main() {
	c := incrementor()
	out := puller(c)
	for v := range out {
		fmt.Println(v)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	var sum int
	go func() {
		for v := range c {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}
