package main

import "fmt"

func makeGreter() func() string {
	return func() string {
		return "Hello Worldd"
	}
}

func main() {
	greeting := makeGreter()
	fmt.Println(greeting())
	fmt.Printf("%T \n", greeting)
}
