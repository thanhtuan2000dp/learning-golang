package main

import "fmt"

type person struct {
	name string
}

func (p person) introduce() {
	fmt.Println("I am", p.name)
}

func main() {
	p1 := person{
		name: "Tuấn",
	}

	p2 := person{
		name: "Tèo",
	}

	p1.introduce()
	p2.introduce()
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
