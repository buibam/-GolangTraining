package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

func main() {
	myname := person{
		first: "Vinh",
		last: "Bui",
		age: 45,
	}
	myname.speak()
}

// Funs
func (p person) speak(){
	fmt.Println("My name: ", p.first, " ", p.last)
	fmt.Println("My age: ", p.age)
}