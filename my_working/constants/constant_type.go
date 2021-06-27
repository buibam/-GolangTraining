package main

import (
	"fmt"
)

const a = 42
const b = 3.14
const c = "This is Vinh"
const(
	d = 14
	e = 6.18
	f = "This is a human."
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}
