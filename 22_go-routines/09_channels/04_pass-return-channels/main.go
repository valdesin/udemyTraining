package main

import "fmt"

func main() {
	c := incrementer()
	out := puller(c)
	for v := range out {
		fmt.Println(v)
	}
}

func incrementer() chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func puller(c chan int) chan int {
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
