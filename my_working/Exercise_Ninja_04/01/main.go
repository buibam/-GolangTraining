package main

import (
	"fmt"
)
func main(){
	var my_array [5]int
	for i := 0; i < 5; i++ {
		my_array[i] = 100 + i
	}
	fmt.Println(my_array)
	fmt.Printf("%T\t\t%v\n", my_array, my_array)

	x := [5]int{73, 23, 75, 84, 99}
	for i, v := range x{
		fmt.Println(i, v)
	}
	fmt.Printf("%T\t\t%v", x, x)
}
