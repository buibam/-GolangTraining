package main

import (
	"fmt"
)
func main(){
	x := 0
	foo(x)
	fmt.Println(x)
	t := 0
	bar(&t)
	fmt.Println(&t)
	fmt.Println(t)
}

// Funcs
func foo(y int){
	fmt.Println(y)
	y = 45
	fmt.Println(y)
}
func bar(z *int){
	fmt.Println(z)
	fmt.Println(*z)
	*z = 75
	fmt.Println(z)
	fmt.Println(*z)
}