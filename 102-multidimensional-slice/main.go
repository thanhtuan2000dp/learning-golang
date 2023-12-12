package main

import "fmt"

func main() {
	slice1 := []string{"test00", "test01", "test02", "test03"}
	slice2 := []string{"test10", "test11", "test12", "test13"}

	fmt.Println(slice1)
	fmt.Println(slice2)

	multiSlice := [][]string{slice1, slice2}
	fmt.Println(multiSlice)
}
