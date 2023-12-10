package main

import (
	"fmt"
)

func main() {
	for x:=0;x<5;x++{
		for y:=0;y<5;y++{
			fmt.Printf("Outer loop is x=%v\t Inner loop is y=%v\n",x,y)
		}
	}
}
