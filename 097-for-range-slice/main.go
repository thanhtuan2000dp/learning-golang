package main

import "fmt"

func main() {
	temp_slice := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}

	for i := 0; i < len(temp_slice); i++ {
		fmt.Printf("Index is %v \t Value is %v \n", i, temp_slice[i])
	}
}
