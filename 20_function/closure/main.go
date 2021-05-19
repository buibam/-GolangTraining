package main

import (
	"fmt"
)
// Scope entire this package
var x int

func main()  {
	// Scope in main() func
	var y int
	fmt.Println(y+5)
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
}

//funcs
func foo() {
	x++
	fmt.Println(x)

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}