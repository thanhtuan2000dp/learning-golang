package main

import "fmt"

func main() {
	temp_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(temp_slice)
	fmt.Println("-------------------------------")

	//[inclusive:exclusive]
	fmt.Println(temp_slice[1:4])
	fmt.Println("-------------------------------")

	//[:exclusive]
	fmt.Println(temp_slice[:4])
	fmt.Println("-------------------------------")

	//[inclusive:]
	fmt.Println(temp_slice[5:])
	fmt.Println("-------------------------------")

	//[:]
	fmt.Println(temp_slice[:])
	fmt.Println("-------------------------------")
}
