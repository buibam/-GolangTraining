package main

import (
	"fmt"
)

type length int
var x length
func main() {

	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)

}
