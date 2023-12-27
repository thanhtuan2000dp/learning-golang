package main

import "fmt"

func main() {
	f := increments()
	fmt.Println(f()) //1
	fmt.Println(f()) //2
	fmt.Println(f()) //3
	fmt.Println(f()) //4
	fmt.Println(f()) //5

	g := increments()
	fmt.Println(g()) //2
	fmt.Println(g()) //1
	fmt.Println(g()) //3
	fmt.Println(g()) //4
	fmt.Println(g()) //5
}

func increments() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
