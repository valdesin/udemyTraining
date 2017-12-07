package main

import "fmt"

func main() {
	c := gen(4, 5)
	out := sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)
}

func gen(n ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range n {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(c chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range c {
			out <- v * v
		}
		close(out)
	}()
	return out
}
