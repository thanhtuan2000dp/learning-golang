package main

import (
	"fmt"
)

func main() {
	temp_slices := []int{42, 43, 44, 45, 46, 47}
	for index, value := range temp_slices {
		fmt.Printf("Index = %v \t Value = %v \n", index, value)
	}
}
