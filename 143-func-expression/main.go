package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("This is anonymous function")
	}

	b := func(a string, b string) {
		fmt.Println("They are", a, "and", b)
	}

	a()
	b("Tuấn", "Tèo")
}
