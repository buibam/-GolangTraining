package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("This should not print.")
	case (2 == 4):
		fmt.Println("This should not print 2.")
	case (3 == 3):
		fmt.Println("Print")
		fallthrough
	case (4 == 4):
		fmt.Println("Also true, should this print?")
		fallthrough
	case (4 == 9):
		fmt.Println("4 == 9, not true")
		fallthrough
	case (11 == 14):
		fmt.Println("11 == 14, not true")
		fallthrough
	case (15 == 15):
		fmt.Println("15 == 15, true")
	default:
		fmt.Println("This is default")
	}
}
