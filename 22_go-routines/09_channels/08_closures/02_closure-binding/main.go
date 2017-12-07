package main

import "fmt"

func main() {
	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, v := range values {
		go func(s string) {
			fmt.Println(s)
			done <- true
		}(v)
	}

	for _ = range values {
		<-done
	}
}
