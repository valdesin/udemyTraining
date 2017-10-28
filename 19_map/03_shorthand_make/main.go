package main

import "fmt"

func main() {
	greeting := make(map[string]string)

	greeting["Tom"] = "Hello"
	greeting["Vlad"] = "Pryvit"

	fmt.Println(greeting)
}
