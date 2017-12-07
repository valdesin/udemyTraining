package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	c1 := sq(in)
	c2 := sq(in)

	for v := range marge(c1, c2) {
		fmt.Println(v)
	}
}

func gen(nums ...int) chan int {
	fmt.Printf("TYPE OF NUMS %T\n", nums)

	out := make(chan int)
	go func() {
		for _, v := range nums {
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

func marge(ch ...chan int) chan int {
	fmt.Printf("TYPE OF CH %T\n", ch)
	var wg sync.WaitGroup

	out := make(chan int)
	wg.Add(len(ch))
	for _, v := range ch {
		go func(c chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
