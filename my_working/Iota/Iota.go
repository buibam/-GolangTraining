package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)
const (
	d = iota
	e
	f
	g
	h
)
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
