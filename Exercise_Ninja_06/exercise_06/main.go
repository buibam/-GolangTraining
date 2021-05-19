package main

import (
	"fmt"
)
func main(){
	// Anomymous
	func(){
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	express := func(){
		fmt.Println("This is a express function")
	}
	express()

	f := func(i int) {
		fmt.Println(i)
	}
	f(50)
}
