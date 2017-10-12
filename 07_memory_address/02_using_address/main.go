package main

import "fmt"

var mettersToYards float64 = 1.09361

func main() {
	var metter float64
	fmt.Println("Enter metter: ")
	fmt.Scan(&metter)
	yards := metter * mettersToYards
	fmt.Println(metter, "is Metters ", yards, "is Yards")
}
