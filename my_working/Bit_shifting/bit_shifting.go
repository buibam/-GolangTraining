package main

import (
	"fmt"
)

const (
	_ = iota
	kb_const = 1 << (iota * 10)
	mb_const = 1 << (iota * 10)
	gb_const = 1 << (iota * 10)
)


func main() {
	x := 12
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	kb := 1024
	mb := 1024*kb
	gb := 1024*mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	//Using iota
	fmt.Printf("%d\t\t\t%b\n", kb_const, kb_const)
	fmt.Printf("%d\t\t\t%b\n", mb_const, mb_const)
	fmt.Printf("%d\t\t%b\n", gb_const, gb_const)

}
