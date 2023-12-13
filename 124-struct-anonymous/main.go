package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Tuấn",
		last:  "Nguyễn",
		age:   23,
	}

	p2 := person{
		first: "Tèo",
		last:  "Phạm",
		age:   19,
	}
	fmt.Printf("%T \n", p1)
	fmt.Printf("%T \n", p2)
}
