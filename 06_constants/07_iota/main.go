package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%b \n", KB)
	fmt.Printf("%d \n", KB)
	fmt.Printf("%b \n", MB)
	fmt.Printf("%d \n", MB)
	fmt.Printf("%b \n", GB)
	fmt.Printf("%d \n", GB)
	fmt.Printf("%b \n", TB)
	fmt.Printf("%d \n", TB)
}
