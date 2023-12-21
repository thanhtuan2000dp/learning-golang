package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.name))
}

func main() {
	p1 := person{
		name: "Tuấn",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer
	p1.writeOut(f)
	p1.writeOut(&b)
	fmt.Println(b.String())
}
