package main

import "fmt"

func main() {
	var recorder [][]string

	student1 := make([]string, 4)
	student1[0] = "Vlad"
	student1[1] = "Fil"
	student1[2] = "180"
	student1[3] = "75"

	recorder = append(recorder, student1)

	student2 := make([]string, 4)
	student2[0] = "Oliver"
	student2[1] = "Stone"
	student2[2] = "185"
	student2[3] = "80"

	recorder = append(recorder, student2)

	fmt.Println(recorder)
}
