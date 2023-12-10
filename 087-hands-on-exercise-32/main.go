package main

import "fmt"

func main() {
	temp_map := map[string]int{
		"Tuan": 42,
		"Teo":  43,
	}

	age := temp_map["Tuan"]
	fmt.Println(age)

	test := temp_map["Q"]
	fmt.Println(test)

	if value, oke := temp_map["Tuan"]; oke {
		fmt.Println(value)
	}
}
