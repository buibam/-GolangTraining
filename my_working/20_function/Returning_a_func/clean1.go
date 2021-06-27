package main

import (
	"fmt"
)
func main(){
	x := foo1()
	fmt.Println(x)
	y := bar()
	fmt.Printf("%T\n", y)
	fmt.Println(y())
}

func foo1() string  {
	return "Hello word"
}

func bar() func() int {
	return func() int {
		return 451
	}
}