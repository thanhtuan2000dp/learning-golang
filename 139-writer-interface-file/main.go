package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s := "Hello gophers!"

	_, err = f.Write([]byte(s))
	if err != nil {
		log.Fatalf("error %s", err)
	}

}