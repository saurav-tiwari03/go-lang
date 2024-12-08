package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	/*
	var a string
	fmt.Print("Enter number A : ");
	  fmt.Scan(&a)
	  fmt.Println("You entered : ",a)
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name : ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is : ",name)
}

