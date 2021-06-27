package main

import (
	"fmt"
)
func main(){
	x := foo()
	fmt.Println(x)

	test := func() int {
		return 451
	}()

	fmt.Println(test)

	y := bar()
	fmt.Printf("%T\n", y)

	i := y()
	fmt.Println(i)
}

func foo() string  {
	s := "Hello word"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}