package main

import (
	"fmt"
)
func main() {
	s := struct {
		last string
		first string
		friends map[string]int
		favDrinks []string
	}{
		last: "Bui",
		first: "Vinh",
		friends: map[string]int{
			"Loc": 916,
			"Ryuchi": 714,
			"Noye": 510,
		},
		favDrinks: []string{
			"water",
			"Beer",
			"wine",
			"soda",
		},
	}
	fmt.Println(s)
	fmt.Println(s.friends)
	for k, val := range s.friends {
		fmt.Printf("%v\t%v\n", k, val)
	}
	for i, val := range s.favDrinks {
		fmt. Printf("%v\t%v\n", i, val)
	}

}