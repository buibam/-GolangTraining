package main

import (
	"fmt"
)

func main() {
	a := 54
	b := 32
	x := a == b
	y := a <= b
	z := a >= b
	t := a != b
	u := a < b
	v := a > b

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(t)
	fmt.Println(u)
	fmt.Println(v)
}
