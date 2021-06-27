package main

import (
	"fmt"
	"encoding/json"
)
type user struct {
	First  string
	Last string
}
func main() {
	p1 := user{
		First : "vinh",
		Last : "bui",
	}
	p2 := user{
		First : "Quynh",
		Last : "Le",
	}
users := []user{p1, p2}
fmt.Println(users)
prJson, err := json.Marshal(users)
if err != nil {
	fmt.Println(err)
}
fmt.Println(string(prJson))
}
