package main
import (
   "fmt"
)
var y = 43
var z = "Shaken, not stirred"
//DECLARE the VARIABLE with the IDENTIFIER "x"
// is of TYPE string
// and I ASSIGN the VALUE "This is a string."
var x string = "This is a string."
// this is a STATIC programing language
//a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC

var a string = `James sadi, "Shaken, not stirred."`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}