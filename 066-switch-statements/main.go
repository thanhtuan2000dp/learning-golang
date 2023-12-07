package main

import "fmt"

func main() {
	//Sequence
	x := 40
	y := 2

	fmt.Printf("x=%v and y=%v\n", x, y)

	// Conditional
	// With range
	switch {
	case x < 40:
		fmt.Println("Less")
	case x == 40:
		fmt.Println("Equal")
	case x > 40:
		fmt.Println("More")
	default:
		fmt.Println("Don't know")
	}

	// With case default
	switch x {
	case 39:
		fmt.Println("x is 39")
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	default:
		fmt.Println("Don't know")
	}

	// With fallthrough

	switch x {
	case 39:
		fmt.Println("x is 39")
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("x is 41")
	default:
		fmt.Println("Don't know")
	}
}
