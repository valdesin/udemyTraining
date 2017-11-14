package main

import (
	"fmt"
	"sort"
)

type people struct {
	Name string
	Age  int
}

func (p people) Strings() {
	fmt.Printf("YAYAYA %s %d", p.Name, p.Age)
}

type ByAge []people

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	person := []people{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(person[0])
	fmt.Println(person)
	sort.Sort(ByAge(person))
	fmt.Println(person)

}
