package main

import "fmt"

func main() {
	v := average(43, 56, 87, 12, 45, 57)
	fmt.Println(v)
}

func average(st ...float64) float64 {
	fmt.Println(st)
	fmt.Printf("%T \n", st)
	var sum float64
	for _, v := range st {
		sum += v
	}
	return sum / float64(len(st))
}
