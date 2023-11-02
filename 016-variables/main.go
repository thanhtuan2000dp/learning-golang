package main

import "fmt"

func main() {
	var a int
	fmt.Println(a)
	a = 1
	fmt.Println(a)

	// Short init
	b := 2
	fmt.Println(b)
	c, d, e, _, f := 2, 4, 5, 6, "hello"
	fmt.Println(c, d, e, f)
}
