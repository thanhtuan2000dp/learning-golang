package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 138", s.String())
}

func main() {
	b1 := book{
		title: "Nhà giả kim",
	}

	var n count = 42

	logInfo(b1)
	logInfo(n)
}
