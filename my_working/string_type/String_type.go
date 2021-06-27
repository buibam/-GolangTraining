package main

import (
	"fmt"
)

func main() {
	s := "This is Vinh."
	t := `"Vinh is I'm"`
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(t)
	fmt.Printf("%T\n", t)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}
	for i, v := range s {
		fmt.Println(i, v)
	}
}
