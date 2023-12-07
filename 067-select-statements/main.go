package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 42
	}()
	select {
	case v1 := <-ch1:
		fmt.Println("Value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("Value from channel 2", v2)
	}
}
