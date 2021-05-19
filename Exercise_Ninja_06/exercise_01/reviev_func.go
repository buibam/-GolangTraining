package main

import (
	"fmt"
)

func main()  {
	fmt.Println(foo())
	fmt.Println(bar())

	n := foo()
	i, s := bar()
	fmt.Println(n, i, s)
}

// Funcs
func foo() int {
	return 1118
}

func bar() (int, string) {
	return 1975, "my birth's year"
}