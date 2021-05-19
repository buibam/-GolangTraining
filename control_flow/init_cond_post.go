package main

import (
	"fmt"
)

func main()  {
	// for init; condition; post {}
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	// for like a while loop
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	// For loop with empty
	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++
	}
	fmt.Print("Done.")
}
