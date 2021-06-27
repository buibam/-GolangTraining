package main

import (
	"fmt"
	"sort"
)
type person struct{
	first string
	age int
}
type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].first < a[j].first }

func main(){
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)
	// There are two ways to sort a slice. First, one can define
	// a set of methods for the slice type, as with ByName, and
	// call sort.Sort. In this first example we use that technique.
	sort.Sort(ByName(people))
	fmt.Println(people)
}
