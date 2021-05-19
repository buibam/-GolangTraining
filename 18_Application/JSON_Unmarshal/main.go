package main

import (
	"encoding/json"
	"fmt"
)
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main(){
	s := `[{"First":"vinhp","Last":"bui","Age":46},{"First":"quynh","Last":"le","Age":35}]`
	bs := []byte(s)
	fmt.Printf("%T\n",s)
	fmt.Printf("%T\n",bs)

	// people := []person{}
	var people []person
	err := json.Unmarshal(bs, &people)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", people)
	for i, v := range people{
		fmt.Println("\nPerson Number ", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
