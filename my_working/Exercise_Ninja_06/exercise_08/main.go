package main

import "fmt"

func main(){
	f := foo()
	fmt.Println(f())
}

// Funcs
func foo() func() int {
	return func() int{

			return 75
		}
}
