package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	value, err := readFile("test.txt")
	if err != nil {
		log.Fatalf("This is error in main Function: %s", err)
	}
	fmt.Println(value)
	fmt.Println(string(value))
}

func readFile(fileName string) ([]byte, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("This is error in ReadFile function: %s", err)
	}
	return content, nil
}
