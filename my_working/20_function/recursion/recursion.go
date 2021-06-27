package main

import (
	"fmt"
)

func main()  {
	n := factorial(4)
	fmt.Println(n)

	byfunc := loopFact(4)
	fmt.Println(byfunc)
}

//Funcs
func factorial(n int) int  {
	if (n == 1){
		return 1
	}
	return n*factorial(n - 1)
}

func loopFact(n int) int {
	result := 1
	for ; n > 1; n-- {
		result *= n
	}
	return result
}