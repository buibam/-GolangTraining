package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {

// Open our jsonFile
jsonFile, err := os.Open("./example.json")
// if we os.Open returns an error then handle it
if err != nil {
    fmt.Println(err)
}
fmt.Println("Successfully Opened example.json")
// defer the closing of our jsonFile so that we can parse it later on
defer jsonFile.Close()

   // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
    tests := make(map[string]interface{})
    
       // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &tests)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    //for i := 0; i < len(tests.Sport); i++ {
    //    fmt.Println("User Type: " + tests)
    //}

    //--------------test

    //fmt.Println("the quiz category types:")
    for key, value := range tests {
            fmt.Println("index : ", key, " value : ", value)
    }
    for key, _ := range tests["quiz"].(map[string]interface{}) {
            fmt.Println("index : ", key)
    }
    //question := make(map[string]interface{})
    if quiz, ok := tests["quiz"].(map[string]interface{}); ok {
        fmt.Printf("quiz = %s\n", quiz["maths"])
        question, _ := quiz["maths"].(map[string]interface{})["q1"].(map[string]interface{})
        fmt.Println(question["question"])
    }
    
} 
