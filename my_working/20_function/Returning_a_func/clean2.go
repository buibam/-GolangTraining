package main

import (
	"fmt"
)
func main(){
	y := bar2()
	fmt.Printf("%T\n", y)
	fmt.Println(bar2()())
}
func bar2() func() int {
	return func() int {
		return 451
	}
}