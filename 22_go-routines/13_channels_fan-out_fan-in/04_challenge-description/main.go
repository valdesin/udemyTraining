package main

import (
	"fmt"
	"time"
)

var publisherId int
var worckerId int

func main() {
	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(1 * time.Millisecond)
}

func publisher(c chan string) {
	publisherId++
	thisId := publisherId
	dataId := 0

	for {
		dataId++
		fmt.Printf("publisher %d is pushing data\n", thisId)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisId, dataId)
		c <- data
	}
}

func workerProcess(in <-chan string) {
	worckerId++
	thisId := worckerId
	for {
		fmt.Printf("%d: waiting for input...\n", worckerId)
		out := <-in
		fmt.Printf("%d: input is %s\n", thisId, out)
	}
}
