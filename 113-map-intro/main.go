package main

import "fmt"

func main() {
	map1 := map[string]int{
		"Tuan": 23,
		"Teo":  19,
	}

	fmt.Println(map1["Tuan"])
	fmt.Printf("%v\n", map1)
	fmt.Println(len(map1))
	fmt.Println("--------------------------------------")

	map2 := make(map[string]int)
	map2["James"] = 19
	map2["Roem"] = 22
	fmt.Println(map2["James"])
	fmt.Printf("%v\n", map2)
	fmt.Println(len(map2))
}
