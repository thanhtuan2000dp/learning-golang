package main

import "fmt"

type person struct {
	first_name         string
	last_name          string
	favorite_ice_cream []string
}

func main() {
	p1 := person{
		first_name:         "Tuấn",
		last_name:          "Nguyễn",
		favorite_ice_cream: []string{"strawberry", "banana"},
	}
	p2 := person{
		first_name:         "Tèo",
		last_name:          "Phạm",
		favorite_ice_cream: []string{"chocolate", "vanilla"},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)

	for _, v := range p1.favorite_ice_cream {
		fmt.Println(p1.first_name, "favorite is", v)
	}

	for _, v := range p2.favorite_ice_cream {
		fmt.Println(p2.first_name, "favorite is", v)
	}
}

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

*/
