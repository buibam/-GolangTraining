package main

import "fmt"

func main() {
	total := foo(2,3,4,5,6,7,8,9)

	fmt.Println("The total is: ", total)
}
// func name ((r receiver) indentifier(parameters)) (returns) {code}
func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, val := range x {
		sum += val
		fmt.Println("for item in index position ", i, "we now add ", val, "to total which is ", sum)
	}
	//fmt.Println("The total is, ", sum)
	return sum
}
