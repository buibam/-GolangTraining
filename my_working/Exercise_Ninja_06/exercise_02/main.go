package main

import (
	"fmt"
)
func main() {
	variadic := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(foo(variadic...))
	ii := bar(variadic)
	fmt.Println(ii)
}

// Funcs
func foo(i ...int) int {
	sum := 0
	for _, val := range i {
		sum += val
	}
	return sum
}

func bar(i []int) int {
	total := 0
	for _, val := range i {
		total += val
	}
	return total
}