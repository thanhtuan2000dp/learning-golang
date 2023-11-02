package main

import "fmt"

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)
const a = 3 >> 1

func main() {
	fmt.Printf("%d \t %b \n", KB, KB)
	fmt.Printf("%d \t %b \n", MB, MB)
	fmt.Printf("%d \t %b \n", GB, GB)
	fmt.Printf("%d \t %b \n", TB, TB)
	fmt.Printf("%d \t %b \n", PB, PB)
	fmt.Printf("%d \t %b \n", EB, EB)

	fmt.Printf("%d \t %b", a, a)
}
