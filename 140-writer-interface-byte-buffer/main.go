package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())

	b.WriteString("Gopher!")
	fmt.Println(b.String())

	b.Reset()
	
	b.Write([]byte("It's new string"))
	fmt.Println(b.String())
}
