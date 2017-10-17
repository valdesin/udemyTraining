package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	changeMe(x)
	fmt.Println(x)
}

func changeMe(v int) {
	fmt.Println(v)
	v = 12
}
