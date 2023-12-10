package main

import "fmt"

func main() {
	temp_map := map[string]int{
		"Tuan": 42,
		"Teo":  43,
	}
	for index, value := range temp_map {
		fmt.Printf("Index is %v \t Value is %v \n", index, value)
	}
}
