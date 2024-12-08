package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var welcome string = "Welcome to the playground!"
	fmt.Println(welcome)
	fmt.Println("Enter your name : ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is : ",name)
}
