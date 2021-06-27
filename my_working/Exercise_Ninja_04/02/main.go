package main

import (
	"fmt"
)
func main(){
	sl := []int{42,43,44,45,46,47,48,49,50,51}
	for i, v := range sl {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\t\t%v", sl, sl)
}
