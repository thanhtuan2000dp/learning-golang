package main

import "fmt"

func main() {
	v1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Tuan",
		friends: map[string]int{
			"Tèo": 23,
			"Tủn": 19,
		},
		favDrinks: []string{"peach tea, mango tea"},
	}

	fmt.Printf("%#v\n", v1)
	fmt.Printf("%v\n", v1)

}

/*
Create a type engine struct, and include this field
electric bool
Create a type vehicle struct, and include these fields
engine
make
model
doors
color
Create two VALUES of TYPE vehicle
use a composite literal
Print out each of these values.
Print out a single field from each of these values.

*/
