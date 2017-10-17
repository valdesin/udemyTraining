package main

import "fmt"

func main() {
	s := make([]string, 1, 25)
	fmt.Println(s)
	changeMe(s)
	fmt.Println(s)
}

func changeMe(s []string) {
	s[0] = "Vlad"
	fmt.Println(s)
}
