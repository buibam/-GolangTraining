package main

import (
	"fmt"
)

type person struct {
	first string
	last string
}
type secretAgent struct {
	person
	ltk bool
}
//func(reciver_name Type) method_name(parameter_list)(return_types){...code...}
func (s secretAgent) speak(){
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "vinh",
			last: "bui",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Quynh",
			last: "le",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
