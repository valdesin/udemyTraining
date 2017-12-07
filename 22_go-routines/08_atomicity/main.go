package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go increment("Foo")
	go increment("Bar")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func increment(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter", atomic.LoadInt64(&counter))
	}
	wg.Done()
}
