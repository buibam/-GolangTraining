package main

import (
	"fmt"
)

func main()  {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tol_sl := sum(slice...)
	fmt.Println("Total all number from 1 to 9: ", tol_sl)

	tol_even := even(sum, slice...)

	fmt.Println("Total even number from 1 to 9: ", tol_even)

	tol_odd := odd(sum, slice...)
	fmt.Println("Total odd number from 1 to 9: ", tol_odd)


}

//Funcs
func sum(x ...int) int  {
	tol := 0
	for _, v := range x {
		tol += v
	}
	return tol
}

func even(f func(x ...int) int, vi ...int) int  {
	var sli []int
	for _, v := range vi {
		if v % 2 == 0 {
			sli = append(sli, v)
		}
	}
	return f(sli...)

}

func odd(f func(x ...int) int, vi ...int) int  {
	var sli []int
	for _, v := range vi {
		if v % 2 != 0 {
			sli = append(sli, v)
		}
	}
	return f(sli...)

}