package main

import "fmt"

func main() {
	c1 := increment("Foo:")
	c2 := increment("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Finel counter:", <-c3+<-c4)
}

func increment(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for v := range c {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}
