package main

import "fmt"

func main() {
	defer testFunc("1")
	defer testFunc("2")
	testFunc("3")
}

func testFunc(s string) {
	fmt.Println(s)
}

/*

“defer” multiple functions in main
show that a deferred func runs after the func containing it exits.
determine the order in which the multiple defer funcs run

*/

// deferred functions run in LIFO order
// last in first out LIFO
