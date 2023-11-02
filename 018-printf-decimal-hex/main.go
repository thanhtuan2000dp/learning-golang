package main

import "fmt"

func main() {
	var a, b, c, d, e, f int = 0, 1, 2, 3, 4, 5
	x, y, z := false, false, true
	fmt.Printf("%v \t %b \t %#X \n", a, a, a)
	fmt.Printf("%v \t %b \t %#X \n", b, b, b)
	fmt.Printf("%v \t %b \t %#X \n", c, c, c)
	fmt.Printf("%v \t %b \t %#X \n", d, d, d)
	fmt.Printf("%v \t %b \t %#X \n", e, e, e)
	fmt.Printf("%v \t %b \t %#X \n", f, f, f)
	fmt.Println(x, y, z)
}
