package main

import (
	"fmt"
)
func main() {
	foo()
	bar("Vinh")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Vinh", "Bui")
	fmt.Println(x)
	fmt.Println(y)
}
// func (r receiver) indentifier(parameters) (return(s)) {...}

func foo() {
	fmt.Println("Hello from foo")
}

// Everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hell0,", s)
}

func woo(st string) string {
	return fmt.Sprintln("Hello from woo, ", st)
}

func mouse (fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a, b
}