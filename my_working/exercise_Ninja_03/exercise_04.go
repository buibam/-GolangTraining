package main

import (
	"fmt"
)

func main() {
	mybd := 1975
	for  {
		if mybd > 2020 {
			break
		}
		fmt.Println(mybd)
		mybd++
	}
}
