package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	avera := average(data...)
	fmt.Println(avera)
}

func average(d ...float64) float64 {
	fmt.Println(d)
	fmt.Printf("%T \n", d)
	total := 0.0
	for _, v := range d {
		total += v
	}
	return total / float64(len(d))
}
