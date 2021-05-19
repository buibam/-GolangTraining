package main
import (
	"fmt"
)
var x int
var g func()
func main(){
	sl_random := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total_even := even(sum, sl_random...)
	fmt.Println("Total string:", total_even)

	f := func(){
		total := 0
		for i := 0; i < 5; i++{
			total += i
			fmt.Println(i)
		}
		fmt.Println(total)
	}
	f()
	fmt.Printf("%T\n",f)
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	g = f
	fmt.Println(g)
	fmt.Printf("%T\n", g)
}

// Funcs
func sum(i ...int) int{
	total := 0
	for _, eval := range i {
		total += eval
	}
	return total
}
func even(f func(i ...int)int, sl ...int) int {
	var sl_even []int
	for _, val := range sl {
		if val%2 == 0 {
			sl_even = append(sl_even, val)
		}
	}
	cal_func := f(sl_even...)
	return cal_func
}