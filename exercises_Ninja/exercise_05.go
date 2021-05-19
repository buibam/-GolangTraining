package main
import(
	"fmt"
)
type number int
var x number
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 100
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}