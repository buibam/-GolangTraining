package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main()  {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	bs := []byte(s)
	var people []person
	err := json.Unmarshal(bs, &people)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(people)
	for index_people, val_people := range people {
		fmt.Println("Person #", index_people)
		fmt.Println("\t",val_people.First, val_people.Last, val_people.Age)
		for _, v := range val_people.Sayings {
			fmt.Println("\t\t",v)
		}
		}
}
