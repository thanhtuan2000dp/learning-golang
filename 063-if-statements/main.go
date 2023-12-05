package main

import "fmt"

func main() {
	//Sequence
	x := 40
	y := 2

	fmt.Printf("x=%v and y=%v\n", x, y)

	//Conditional
	if x < 42 {
		fmt.Println("Less")
	} else if x == 42 {
		fmt.Println("Equal")
	} else {
		fmt.Println("More")
	}

}
