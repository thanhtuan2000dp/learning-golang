package main

import "fmt"

var x int = 2

const y string = "hello"

func main() {
	z := 2.5
	fmt.Printf("Types: %T \tValue:%v\n", x, x)
	fmt.Printf("Types: %T \tValue:%v\n", y, y)
	fmt.Printf("Types: %T \tValue:%v\n", z, z)
}
