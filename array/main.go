package main

import "fmt"

func main() {
	var fruitList = [4]string{"Apple", "Banana", "Orange", "Mango"}
	fmt.Println("Fruit list:", fruitList)
	fmt.Printf("Length: %d, Type: %T\n", len(fruitList), fruitList)
}
