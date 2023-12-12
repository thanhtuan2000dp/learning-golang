package main

import "fmt"

func main() {
	// Params: First is Type, Second is Length, Third is capacity
	temp_slice := make([]int, 0, 10)
	fmt.Println(temp_slice)
	fmt.Println(len(temp_slice))
	fmt.Println(cap(temp_slice))
}
