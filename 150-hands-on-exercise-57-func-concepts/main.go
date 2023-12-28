package main

import "fmt"

func main() {

	x := sum([]int{1, 2, 3, 4, 5})
	fmt.Println(x)
}

func sum(ii []int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

//Don't like name returns
/*
func sum(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}
*/
