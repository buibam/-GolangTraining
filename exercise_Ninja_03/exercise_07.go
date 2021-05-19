package main

import (
	"fmt"
)

func main() {
	age := 91
	if age < 18 {
		fmt.Println("You are not enough age to vote")
	} else if age > 80 {
		fmt.Println("You should vote by mail.")
	} else {
		fmt.Println("You can vote by any type of vote if you feel confortable.")
	}
}
