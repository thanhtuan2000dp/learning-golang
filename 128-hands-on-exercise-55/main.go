package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors string
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{electric: true},
		make:   "car",
		model:  "Honda City",
		doors:  "4",
		color:  "White",
	}

	v2 := vehicle{
		engine: engine{electric: false},
		make:   "bike",
		model:  "Martin",
		doors:  "0",
		color:  "deep blue",
	}

	fmt.Printf("%#v\n", v1)
	fmt.Printf("%#v\n", v2)

	println(v1.electric, v1.make, v1.model, v1.doors, v1.color)
	println(v2.electric, v2.make, v2.model, v2.doors, v2.color)
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
