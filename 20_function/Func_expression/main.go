package main

import (
	"fmt"
)
func main(){
	f := func(){
		fmt.Println("my first expression")
	}
	f()
	g := func(x int){
		fmt.Println("x's value is:", x)
	}
	g(45)

}
