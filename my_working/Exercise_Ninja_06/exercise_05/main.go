package main
import (
	"fmt"
	"math"
)
type circle struct {
	r float64
}
type square struct {
	width float64
}

type shape interface {
	area() float64
}

func main() {
	c := circle{
		r: 45.3,
	}
	sq := square{
		width: 22.556,
	}
	info(c)
	info(sq)

}

//Funcs
func (c circle)area() float64{
	return  c.r * c.r * math.Pi
}
func (r square )area() float64{
	return r.width*r.width
}

func info(s shape){
	fmt.Println(s.area())
}

