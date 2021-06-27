package main

import (
	"fmt"
)
func main(){
	defer defer1()
	normal1()
	normal2()
	defer defer2()
}
// Functions
func defer1() {
	fmt.Println("Run defer 1.")
}
func defer2() {
	fmt.Println("Running defer 2.")
}
func normal1() {
	fmt.Println("Normal1 running.")
}
func normal2() {
	fmt.Println("Normal2 running.")
}
