package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	ice_cream_flavors []string
}
func main() {
	p1 := person{
		first: "Vinh",
		last: "Bui",
		ice_cream_flavors: []string{
							"coconut",
							"vanilla",
							"mango",
						},
	}
	p2 := person{
		first: "James",
		last: "Bone",
		ice_cream_flavors: []string {
					"vanilla",
					"strawberry",
					"capuccino",
					},
						}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, val := range p1.ice_cream_flavors {
		fmt.Printf("\t%v\t%v\n", i, val)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, val := range p2.ice_cream_flavors {
		fmt.Printf("\t%v\t%v\n",i, val)
	}
}
