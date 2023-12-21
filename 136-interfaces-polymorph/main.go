package main

import "fmt"

type person struct {
	name string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) introduce() {
	fmt.Println("I am", p.name)
}

func (sa secretAgent) introduce() {
	fmt.Println("I am a secret agent", sa.name)
}

type human interface {
	introduce()
}

func saySomething(h human) {
	h.introduce()
}

func main() {
	sa1 := secretAgent{
		person: person{
			name: "Tuấn",
		},
		ltk: true,
	}

	p2 := person{
		name: "Tèo",
	}

	saySomething(sa1)
	saySomething(p2)
}
