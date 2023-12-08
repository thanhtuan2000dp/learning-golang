package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x, y := rand.Intn(10), rand.Intn(10)
	fmt.Printf("Value x=%v\ty=%v\n", x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("Both less than 4")
	case x > 6 && y > 6:
		fmt.Println("Both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("X greater than 4 and less than 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("None of the previous were met")
	}

}
