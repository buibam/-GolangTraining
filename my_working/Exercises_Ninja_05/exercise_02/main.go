package main

import (
	"fmt"
)
type vehicle struct {
	doors	int
	color	string
}
type struck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury	bool
}
func main() {
	toyota_truck := struck {
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourwheel: true,
	}
	honda_sedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "golden",
		},
		luxury: false,
	}
	fmt.Println(toyota_truck)
	fmt.Println(toyota_truck.doors)
	fmt.Println(toyota_truck.color)
	fmt.Println(toyota_truck.vehicle)
	fmt.Println(honda_sedan)
}
