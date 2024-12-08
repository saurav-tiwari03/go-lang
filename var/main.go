package main 

import "fmt"

const Token string = "1234567890"

func main () {
	var name string = "John"
	age1 := 21;
	fmt.Println("Variables in Go : ",name,age1,Token)
	fmt.Printf("Variables in Go : %T\n", name)
}