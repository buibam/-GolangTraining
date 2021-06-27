package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

func main(){

	Displaywelcom()
	var filename string = "test-go2.txt"
	var inputText string = "Hello, This is my first Go program by using function."
	//Open a file name "test-go.txt"
	err := OpenWriteToFile(filename, inputText)
    if err != nil {
        fmt.Println("Erro while open the test-go2.txt file", err)
    }

	// Check content of the file
	CheckaFile(filename, inputText)
}

//--------------------------------------------

//Function code
func Displaywelcom(){
	fmt.Println("Welcome to my Go Program")
}

func OpenWriteToFile(filename string, inputstring string) error {
    file, err := os.Create(filename)
    if err != nil {
    		fmt.Printf("\nOpen the file failed: %s\n", err)
        return err
    }
    defer file.Close()

    _, err = io.WriteString(file, inputstring)
    if err != nil {
    	fmt.Printf("\nfailed writing to file: %s\n", err)
        return err
    }
    return file.Sync()
}

func CheckaFile(filename string, myStr string){
	rd, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("\nfailed reading data from file: %s\n", err)
	}else {

	compareStr := string(rd) == myStr

	if !compareStr {
			fmt.Println("\nThe file doesn't containt my text.")
		} else {
			fmt.Println("\nMy text has been added to test-go2.txt.")
		}	
	}
	//Print the content of my file
	fmt.Println("\nThe content of test-go.txt is.")
	fmt.Println("\n",string(rd))
}