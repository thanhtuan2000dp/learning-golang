package main

import "fmt"

func main() {
	temp_slice := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	fmt.Println(temp_slice, cap(temp_slice), len(temp_slice))
	fmt.Println("--------------------------------")
	temp_slice = append(temp_slice, "Hello", "World")
	fmt.Println(temp_slice, cap(temp_slice), len(temp_slice))
	fmt.Println("--------------------------------")

}
