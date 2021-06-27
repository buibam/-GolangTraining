package main

import (
	"fmt"
)
func main(){
	defer calDefer()
	slice := []int {1,2,3,4,5,6,7}
	total := foo(slice...)
	fmt.Println(total)

}

// Funcs
func foo(i ...int) int{
	sum := 0
	for _, val := range i {
		sum += val
	}
	return sum
}

func calDefer(){
	fmt.Println("Using defer to call calDefer()")
}