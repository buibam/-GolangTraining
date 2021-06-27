package main

import (
	"fmt"
	"io"
	"os"
)
func main(){
	fmt.Println("Hello, I'm here.")
	fmt.Fprintln(os.Stdout, "Hello, this is os.Stdout")
	io.WriteString(os.Stdout, "Hello, this is WriteString")
}
