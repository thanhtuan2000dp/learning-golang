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

	list_persons := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for key, value := range list_persons {
		fmt.Println(key)

		for _, v := range value.favorite_ice_cream {
			fmt.Println(value.first_name, "favorite is", v)
		}
	}
}

/*
Take the code from the previous exercise, then store the VALUES of type person in a map with the KEY of last name. Access each value in the map. Print out the values, ranging over the slice.
*/
