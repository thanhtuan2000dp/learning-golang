package main

import "fmt"

func main() {
	var a int
	fmt.Println(a)

	b := 3
	fmt.Println(b)

	c, d := 4, 5
	fmt.Println(c, d)

	var e float64 = 4.56
	fmt.Printf("%v\t%T\n", e, e)

	f, g, _ := 7, 8, 9
	fmt.Println(f, g)
}
