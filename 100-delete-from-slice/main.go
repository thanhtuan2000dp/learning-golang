package main

import "fmt"

func main() {
	temp_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(temp_slice)
	fmt.Println("-------------------------------")

	temp_slice = append(temp_slice[:4], temp_slice[5:]...)
	fmt.Println(temp_slice)
	fmt.Println("-------------------------------")
}
