package main

import "fmt"

func main() {
	for v := range sq(sq(gen(3, 4))) {
		fmt.Println(v)
	}

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
