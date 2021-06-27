package main

import (
	"fmt"
)
func main(){
	m := map[string] []string {`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`}, `no_dr`: []string{`Being evil`, `Ice cream`, `Sunsets`}}
	fmt.Println(m)
	m["fleming_ian"] = []string{"steak", "cigars", "enpionage"}
	for k, v := range m {
		fmt.Println("This is the record for: ", k)
		for i, val := range v {
			fmt.Printf("%v:\t%v\n", i, val)
		}
	}
	delete(m, "fleming_ian")
	for outIndex, outRecord := range m {
		fmt.Println("This is the record for: ", outIndex)
		for inIndex, inValue := range outRecord {
			fmt.Printf("%v \t%v\n", inIndex, inValue)
		}
	}

}
