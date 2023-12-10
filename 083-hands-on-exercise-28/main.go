package main

import (
	"fmt"
)

func main() {
	x:=50
	for i:=0;i<x;i++ {
		if (i%2==0){
			continue
		}
		fmt.Printf("ODD number is %v\n",i)
	}
}
