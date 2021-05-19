package main

import (
	"fmt"
)

const (
	a = 15

	b = 3.14
	c = true
	d = "This a string."
)
const (
	x int = 42
	y float32 = 3.14
	z string = "This is another string."
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}