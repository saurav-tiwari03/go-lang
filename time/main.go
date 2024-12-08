package main

import (
	"fmt"
	"time"
)

func main() {
	formatedTime := time.Now().Format("01-02-2006 15:04:05 Monday")
	fmt.Println("Time : ", time.Now())
	fmt.Println("Formated Time : ", formatedTime)
}
