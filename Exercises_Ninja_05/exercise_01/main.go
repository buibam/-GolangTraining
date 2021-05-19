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
	m := map [string] person {
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(m)
	for k, val := range m {
		fmt.Println(k)
		fmt.Println(val.first)
		fmt.Println(val.last)
		for i, v := range val.ice_cream_flavors {
			fmt.Println(i, v)
		}
	}
}
