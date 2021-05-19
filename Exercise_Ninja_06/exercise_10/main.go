package main

import (
	"fmt"
)
func main(){
	fmt.Println(foo()())
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	g := foo()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

	j := bar()
	fmt.Println(j)
	fmt.Println(j)
	fmt.Println(j)
	fmt.Println(j)
}

// Funcs
func foo() func()int{
	x := 0
	return func() int {
			x ++
		return x
	}
}

func bar() int{
	y := 0
	y++
	func(){
		fmt.Println(y)
	}()

	return y
}
