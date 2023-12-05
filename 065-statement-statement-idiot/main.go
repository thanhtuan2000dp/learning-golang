package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Sequence
	x := 40
	y := 2

	fmt.Printf("x=%v and y=%v\n", x, y)

	//Conditional
	if z := 2 * rand.Intn(40); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)

	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}
}
