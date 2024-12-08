package main

import "fmt"
//Automatically resizes the array

func main()  {
	var fruitList = []string{"Apple", "Banana", "Orange", "Mango"}
	fmt.Println("Fruit list:", fruitList)
	fmt.Printf("Length: %d, Type: %T\n", len(fruitList), fruitList)

	fruitList = append(fruitList, "Peach", "Grapes")
	fmt.Println("Fruit list:", fruitList)

	fruitList = append(fruitList[1:3]) //[1:] means from 1 to end it keeps the items from 1 to 3
	fmt.Println("Fruit list:", fruitList)

	//make function
}