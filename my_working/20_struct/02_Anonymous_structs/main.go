package main

import (
	"fmt"
)
/* type person struct {
	first	string
	last	string
	age  	int
}*/
func main() {
	//Anonymous
	p1 := struct {
		first	string
		last	string
		age  	int
	}{
		first: "Vinh",
		last: "Bui",
		age: 45,
	}
	fmt.Println(p1)

}
