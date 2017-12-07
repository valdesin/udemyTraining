package main

import (
	"fmt"
	"sync"
	//"time"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()
	//time.Sleep(time.Millisecond * 3)
	for v := range c {
		fmt.Println(v)
	}
}
