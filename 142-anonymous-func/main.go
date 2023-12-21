package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("This is anonymous function")
	}()

	func(a string, b string) {
		fmt.Println("They are", a, "and", b)
	}("Tuấn", "Tèo")
}
