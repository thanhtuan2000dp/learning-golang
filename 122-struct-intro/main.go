package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Tuấn",
		last:  "Nguyễn",
		age:   23,
	}

	p2 := person{
		first: "Tèo",
		last:  "Phạm",
		age:   19,
	}

	fmt.Println(p1)
	fmt.Printf("%T \t %#v", p2, p2)
}
