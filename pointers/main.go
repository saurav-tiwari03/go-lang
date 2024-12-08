package main

import "fmt"

func main() {
	// var ptr *int
	var myNumber int = 23
	var ptr = &myNumber
	fmt.Println("Value of ptr : ", ptr)
	fmt.Println("Value of myNumber : ", *ptr)
}