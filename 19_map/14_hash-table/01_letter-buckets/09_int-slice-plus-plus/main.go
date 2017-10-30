package main

import "fmt"

func main() {
	backets := make([]int, 1)
	fmt.Println(backets)
	backets[0] = 42
	fmt.Println(backets)
	backets[0]++
	fmt.Println(backets)
}
