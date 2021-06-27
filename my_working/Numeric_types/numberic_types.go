package main

import (
	"fmt"
)
var a int
var b float64
func main() {
	x := 42
	y := 42.34253
	a = 123
	b = 3.14
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}
