package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		j := i%4
		fmt.Printf("%v/4 remain %v.\n", i, j)
	}
}
