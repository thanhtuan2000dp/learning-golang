package main

import "fmt"

func main() {
	x := 40
	y := 5
	println("---------------- Loop 1 -------------------------")
	for i := 0; i < 5; i++ {
		fmt.Printf("Value of i is:\t%v\n", i)
	}

	println("---------------- Loop 2 -------------------------")
	for y < 10 {
		fmt.Printf("Valye of y is:\t%v\n", y)
		y++
	}

	println("---------------- Loop 3 -------------------------")
	for {
		if y > 20 {
			break
		}
		fmt.Printf("Value of y is:\t%v\n", y)
		y++
	}

	println("---------------- Loop 4 -------------------------")
	for i := 0; i < x; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("Even number is:\t%v\n", i)
	}

}
