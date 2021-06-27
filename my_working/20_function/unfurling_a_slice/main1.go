package main

import (
"fmt"
)
func main() {
	x := sum("jame", 4, 5, 6, 7)
	// we can pass none parameter
	// x := sum ()
	fmt.Println("The total is:", x)
}

func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in dex position ", i, "we are now adding ", v, "to total which is \", sum)
	}
	return sum

}