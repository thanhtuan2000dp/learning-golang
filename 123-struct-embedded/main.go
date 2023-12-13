package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person // Embedded struct
	ltk    bool
}

func main() {
	p1 := secretAgent{
		person: person{
			first: "Tuấn",
			last:  "Nguyễn",
			age:   23,
		},
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Println(p1.first, p1.last, p1.age, p1.ltk, p1.person)
}
