package main

import "fmt"

func main() {
	fmt.Println(paradise("New york"))
}

func paradise(location string) string {
	return fmt.Sprint("My idea of paradise is ", location)
}
