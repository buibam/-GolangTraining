package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main(){
	//Check the file exist or not - working now
	if _, err := os.Stat("test-go2.txt"); os.IsNotExist(err) {
		fmt.Println("File definitely does not exist.")
	}
	//--------------------------------------

	//Open a file name "test-go.txt"
	fmt.Println("Welcome to my Go Program")
	f, err := os.Create("C:\\code\\test-go.txt")
	if err != nil {
		fmt.Printf("\nfailed creating file: %s\n", err)
		return
	} else {
		fmt.Println("The test-go.txt file is created successfully.")
	}

	//Write a string into test-go.txt
	len, err := f.WriteString("Hello, This is my first Go program.")
	if err != nil {
		fmt.Printf("\nfailed writing to file: %s\n", err)
	} else {
		fmt.Printf("\nLength: %d bytes\n", len)
    	fmt.Printf("\nFile Name: %s\n", f.Name())
	}
	f.Close()
	// Check content of the file
	rd, err := ioutil.ReadFile("test-go.txt1111")
	if err != nil {
		fmt.Printf("Can't read content from test-go.txt: %s\n", err)
		fmt.Println("Can't read content from test-go.txt:", err)
		fmt.Println("Recieved error",err,"content from test-go.txt")
		return
	}else {

		compareStr := string(rd) == "Hello, This is my first Go program."
		if !compareStr {
			fmt.Println("\nThe file doesn't containt my text.")
		} else {
			fmt.Println("\nMy text has been added to test-go.txt.")
		}	
	}

	//Print the content of my file
	fmt.Println("\nThe content of test-go.txt is.")
	fmt.Println("\n",string(rd))

}