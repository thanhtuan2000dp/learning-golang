package main

import "fmt"

func main() {
	slice1 := []string{"James", "Bond", "Shaken, not stirred"}
	slice2 := []string{"Miss", "Moneypenny", "I'm 008."}

	multiSlice := [][]string{slice1, slice2}

	for indexRecord, valRecord := range multiSlice {
		fmt.Printf("Record index is %v \n", indexRecord)
		for indexData, valData := range valRecord {
			fmt.Printf("Data index is %v \t Value is %v \n", indexData, valData)
		}
		fmt.Println("-------------------------------------")
	}
}

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."

Range over the records, then range over the data in each record.

*/
