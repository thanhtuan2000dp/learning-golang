package main

import "fmt"

func main() {
	map1 := map[string]int{
		"Tuan": 23,
		"Teo":  19,
	}
	map1["James"] = 22
	fmt.Printf("%v\n", map1)

	fmt.Println("---------- Delete --------------")
	delete(map1, "Teo")
	if v, ok := map1["Teo"]; !ok {
		fmt.Println("Key didn't exist!")
	} else {
		fmt.Printf("Value is %v\n", v)
	}

}
