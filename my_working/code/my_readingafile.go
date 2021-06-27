package main

import (
	"fmt"
	"io/ioutil"
)

func main(){

	Displaywelcom()
	var filename string = "./example.json"
	// Read content of the file
	readafile(filename)
	readJSON(filename)
}

//--------------------------------------------

//Function code
func Displaywelcom(){
	fmt.Println("Welcome to my Go Program")
}

func readafile(filename string){
	rd, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("\nfailed reading data from file: %s\n", err)
	}else {
	//Print the content of my file
	fmt.Println("\nThe content of example.json is.")
	fmt.Println("\n",string(rd))
	}
}

func readJSON(fileName string, filter func(map[string]interface{}) bool){
	datas := []map[string]interface{}{}

	file, _ := ioutil.ReadFile(fileName)
	json.Unmarshal(file, &datas)

	filteredData := []map[string]interface{}{}

	for _, data := range datas {
		// Do some filtering
		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	fmt.Println(filteredData)
}