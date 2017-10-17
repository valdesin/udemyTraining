package main

import "fmt"

func makeGreter() func() string {
	return func() string {
		return "Hello World"
	}
}

func main() {
	greeting := makeGreter()
	fmt.Println(greeting())
}
