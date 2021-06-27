package main

import (
	"fmt"
)
func main(){
	x := 75
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)

	// var y *int = &x
	y := &x
	fmt.Println(y)

	fmt.Printf("%T\n",y)
	fmt.Println(*y)		// * gives you the value stored in the address
	fmt.Println(*&x)

	*y = 43
	fmt.Println(x)

}
