package main

import (
	"fmt"
)

const (
	_ = iota
	year1 = 2016 + iota
	year2 = 2016 + iota
	year3 = 2016 + iota
	year4 = 2016 + iota
)
func main(){

	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)

}
