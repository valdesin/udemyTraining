package main

import "fmt"

func main() {
	m := map[int]int{}
	m[2] = 23
	m[82] = 44
	m[54] = 123

	ch := make(chan int)

	ch2 := make(chan int)
	go func() {
		var i int
		for v := range ch2 {
			i += v
			if i == len(m) {
				close(ch)
			}
		}
	}()

	go func() {
		for _, v := range m {
			ch <- v
			ch2 <- 1
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}

	close(ch2)
}
