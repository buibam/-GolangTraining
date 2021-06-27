package main

import (
	"fmt"
)

func main() {
	name := "Bond"
	switch name {
	case "Moneypenny", "My bank", "Bond":
		fmt.Println("Moneypenny or Oh no or Bond")
	case "Oh no":
		fmt.Println("Oh no")
	case "No Bond":
		fmt.Println("Bond James")
	case "Q":
		fmt.Println("This is Q")
	default:
		fmt.Println("This is deufault.")


	}
}