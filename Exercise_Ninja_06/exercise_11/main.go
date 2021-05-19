package main

import (
	"fmt"
)
func main(){
	y := countint(5)
	fmt.Println(y)
}

// Funcs
func countint(i int) int{
	if i == 1 {
		return i
	}
	fmt.Println(i)
	return countint(i-1)
}