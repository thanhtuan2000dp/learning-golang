package main

import "fmt"

func main() {
	foo()

	bar("Tuấn")

	v3 := aloha("Aha")
	fmt.Println(v3)

	v4, age := petYear("Bobby", 3)
	fmt.Println(v4, age)
}

func foo() {
	fmt.Println("Hello foo func")
}

func bar(name string) {
	fmt.Printf("Hello my name is %v\n", name)
}

func aloha(name string) string {
	return fmt.Sprint(name, " has 5 year old")
}

func petYear(name string, age int) (string, int) {
	return fmt.Sprint(name, " has age is"), age
}

/*

func syntax

no params, no returns
1 param, no returns
1 param, 1 return
2 params, 2 returns
*/
