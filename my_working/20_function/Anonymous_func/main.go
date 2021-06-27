package main

import (
	"fmt"
)
func main(){
	foo()
	func(){
		fmt.Println("Anomymous func ran")
	}()
	func(x int){
		fmt.Println("Value of of x: ", x)
	}(42)
}

func foo()  {
	fmt.Println("foo ran")
}