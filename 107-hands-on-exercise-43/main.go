package main

import "fmt"

func main() {
	arr := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, value := range arr {
		fmt.Printf("Type is %T \t Value is %v \n", value, value)
	}
}
