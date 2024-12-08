package main

import "fmt"

func main() {
	fmt.Println("Welcome to the playground!")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Orange"
	fruitList[3] = "Mango"
	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Length of fruitList is : ", len(fruitList))
	fmt.Printf("Type : %T",fruitList[0])
}
