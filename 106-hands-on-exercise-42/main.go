package main

import "fmt"

func main() {
	arr := [5]int{2, 3, 1, 4, 5}

	for _, value := range arr {
		fmt.Printf("Type is %T \t Value is %v \n", value, value)
	}
}
