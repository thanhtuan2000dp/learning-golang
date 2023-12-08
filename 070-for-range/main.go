package main

import "fmt"

func main() {
	//For range loop
	//Data structures - slice
	s1 := []int{41, 42, 43, 44, 45, 46}
	for index, val := range s1 {
		fmt.Println("ranging over a slice", index, val)
	}

	//For range loop
	//Data structures - map
	m := map[string]int{
		"Tuan": 23,
		"Teo":  20,
	}
	for index, val := range m {
		fmt.Println("ranging over a map", index, val)
	}
}
