package main

import (
	"fmt"
)
func main(){
	g := func(ii []int) int{
		total := 0
		for _, value := range ii{
			total += value
		}
		return total
	}
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(g, xi...))
}

// Func
// This func will pass a func in its parameter
func foo(f func(i []int) int, ii ...int) int{
	n := f(ii)
	n++
	return n
}