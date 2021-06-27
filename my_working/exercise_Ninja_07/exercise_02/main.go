package main

import (
	"fmt"
)

type person struct {
	name string
}

func main(){
	p1 := person{
		name: "Vinh Bui",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

// Func
func changeMe(p *person) {
	p.name = "Miss Moneypenny"
	// (*p).first = "Miss Moneyp"
}