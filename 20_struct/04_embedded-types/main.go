package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			first: "Vlad",
			last:  "Fill",
			age:   23,
		},
		LicenseToKill: false,
	}

	p2 := doubleZero{
		person: person{
			first: "Todd",
			last:  "Wilyam",
			age:   44,
		},
		LicenseToKill: true,
	}

	fmt.Println(p1.first, p1.last, p1.age, p1.LicenseToKill)
	fmt.Println(p2.first, p2.last, p2.age, p2.LicenseToKill)
}
