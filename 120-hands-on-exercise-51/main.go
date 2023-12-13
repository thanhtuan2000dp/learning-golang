package main

import "fmt"

func main() {
	tempMap := map[string][]string{
		"bond_james":       {`shaken, not stirred`, `martinis`, `fast cars`},
		"moneypenny_jenny": {`intelligence`, `literature`, `computer science`},
		"no_dr":            {`cats`, `ice cream`, `sunsets`},
	}

	tempMap["fleming_ian"] = []string{`steak`, `cigars`, `espionage`}

	delete(tempMap, "no_dr")
	for keyRecord, valRecord := range tempMap {
		fmt.Printf("----------- %v ------------\n", keyRecord)
		for indexData, valData := range valRecord {
			fmt.Printf("Index = %v \t Value = %v \n", indexData, valData)
		}
	}
}
