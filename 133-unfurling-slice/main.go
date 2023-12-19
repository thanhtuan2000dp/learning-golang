package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	s := sum(x...)
	fmt.Print("Sum is ", s)
}

func sum(params ...int) int {
	fmt.Println(params)
	fmt.Printf("%T\n", params)

	n := 0
	for _, v := range params {
		n += v
	}
	return n
}
