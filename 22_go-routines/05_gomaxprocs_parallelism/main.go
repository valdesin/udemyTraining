package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 22; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Millisecond * 5)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 22; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Millisecond * 20)
	}
	wg.Done()
}
