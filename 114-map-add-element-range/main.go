package main

import "fmt"

func main() {
	map1 := map[string]int{
		"Tuan": 23,
		"Teo":  19,
	}
	map1["James"] = 22

	fmt.Printf("%v\n", map1)

	for key, val := range map1 {
		fmt.Printf("Key is %v \t Value is %v \n", key, val)
	}
	fmt.Println("-----------------------------------")
	for _, val := range map1 {
		fmt.Printf("Value is %v \n", val)
	}
	fmt.Println("-----------------------------------")
	for key := range map1 {
		fmt.Printf("Key is %v \n", key)
	}
}
