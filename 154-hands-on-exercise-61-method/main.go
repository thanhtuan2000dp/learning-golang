package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, "and I'm", p.age, "years old")
}

func main() {
	p1 := person{
		first: "Tuấn",
		age:   23,
	}

	p1.speak()
}

/*
Create a user defined struct with
the identifier “person”
the fields:
first
age
attach a method to type person with
the identifier “speak”
the method should have the person say their name and age
create a value of type person
call the method from the value of type person
*/
