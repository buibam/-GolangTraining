package main

import (
	"fmt"
	"math"
)
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func main()  {
	// Call direct area() function
	cir := circle{
		radius: 5,
	}
	fmt.Println(cir.area())

	// Call via shap interface
	cir_interf := circle{
		radius: 7,
	}
	// s := cir_interf.area()
	//fmt.Println("The area of circle with radius = 7 is: ", s)
	fmt.Println("The area of circle with radius = 7 is: ", cir_interf.area())

	cir_radius := circle{
		radius: 10,
	}
	info(cir_radius)

	// Pointer
	cir_pointer := circle{
		radius: 10,
	}
	fmt.Printf("%T\n", &cir_pointer)
	info(&cir_pointer)
}
// Funcs
func (c circle) area() float64{
	return c.radius * c.radius * math.Pi
}

func info(s shape){
	fmt.Println("The area of circle is: ", s.area())
}