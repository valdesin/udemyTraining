package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := incrementor(2)

	var count int

	for v := range c {
		count++
		fmt.Println(v)
	}
	fmt.Println("Final Count: ", count)
}

func incrementor(n int) chan string {
	out := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				out <- fmt.Sprintln("Process: "+strconv.Itoa(i)+" printing: ", k)
			}
			done <- true
		}(i)
	}
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(out)
	}()
	return out
}
