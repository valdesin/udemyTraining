package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Make"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(s string) chan string {
	out := make(chan string)
	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%s : %d", s, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return out
}

func fanIn(input1, input2 chan string) chan string {
	out := make(chan string)
	go func() {
		for v := range input1 {
			out <- v
		}
	}()

	go func() {
		for v := range input2 {
			out <- v
		}
	}()
	return out
}
