package main

import (
	"fmt"
)

func main() {
	for i := 33; i < 123; i++ {
		fmt.Printf("%v\t\t%#x\t\t%#U\n", i, i, i)
	}
}
