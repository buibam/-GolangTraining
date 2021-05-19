package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)
// Tests struct which contains
// an array of quizs
type Quizs struct {
    Quizs Quiz `json:"quiz"`
}

// Sports struct which contains a array of sport question
type Quiz struct {
    Sports Sport `json:"sport"`
    Maths Math `json:"maths"`
}

// QuestionNum struct which contains parts of a question in the quiz

type Sport struct {
    Q1  Qs1 `json:"q1"`
}
type Math struct {
    Q1  Qm1 `json:"q1"`
    Q2  Qm2 `json:"q2"`
}

type Qs1 struct {
    Question  string `json:"question"`
    Options   []string `json:"options"`
    Answer    string `json:"answer"`
}

type Qm1 struct {
    Question  string `json:"question"`
    Options   []string `json:"options"`
    Answer    string `json:"answer"`
}

type Qm2 struct {
    Question  string `json:"question"`
    Options   []string `json:"options"`
    Answer    string `json:"answer"`
}

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
    var tests Quizs

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &tests)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    //for i := 0; i < len(tests.Sport); i++ {
    //    fmt.Println("User Type: " + tests)
    //}

    //test

    fmt.Println("the quiz category types:")
    fmt.Println(tests.Quizs)
    fmt.Println("all the questions:")
    //for i := 0; i < len(tests.Sport); i++ {
    //    fmt.Println(tests.Quizs.Maths.Q1.Question)
    //}
    fmt.Println(tests.Quizs.Maths.Q1.Question)
} 