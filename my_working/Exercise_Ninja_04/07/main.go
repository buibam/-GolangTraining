package main

import (
	"fmt"
)
func main(){
	mystring := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	fmt.Println(mystring)
	for _, x := range mystring {
		for _, y := range x {
			fmt.Println(y)
		}
	}
	// another way
	fmt.Println("Another way.")
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)
	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)
	for ixxs, xs := range xxs {
		fmt.Println("Record: ", ixxs)
		for subindex, v := range xs {
			fmt.Printf("\t index position: %v\t value: %v\n", subindex, v)

		}
	}
}