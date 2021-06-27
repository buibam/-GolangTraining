package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)
//------------global varible---------
var jsonFile string 
//------------functions--------------
func readjsonFile(fileName string) map[string]interface{}{
	// Open our jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
    	fmt.Println(err)
	}
	fmt.Println("Successfully Opened example.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

   //Read the jsonFile and stor in byteValue
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // Variable to store json file in map
        jsonInmap := make(map[string]interface{})
    
    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'quiz' which we defined above
    json.Unmarshal(byteValue, &jsonInmap)

    return jsonInmap
}

func displayjson(inputdata map[string]interface{}) {
	//Display the quiz category
	for key, _ := range inputdata["quiz"].(map[string]interface{}) {
        fmt.Println("quiz : ", key)
    }

    //Display all question
    fmt.Println("Sport questions:" )
    question, _ := inputdata["quiz"].(map[string]interface{})["sport"].(map[string]interface{})["q1"].(map[string]interface{})
    fmt.Println(question["question"])
    fmt.Println("Maths questions:" )
    questionMath1, _ := inputdata["quiz"].(map[string]interface{})["maths"].(map[string]interface{})["q1"].(map[string]interface{})
    fmt.Println(questionMath1["question"])
    questionMath2, _ := inputdata["quiz"].(map[string]interface{})["maths"].(map[string]interface{})["q2"].(map[string]interface{})
    fmt.Println(questionMath2["question"])
}

func main() {
	jsonFile = "./example.json"
	//Read a json file
	readingjs := readjsonFile(jsonFile)

	// Display Json conten whatever we want
	displayjson(readingjs)
}