package main

import (
	"net/http"
	"fmt"
	"golang/todo-list"
)

func main()  {
	r := router.Router()
	fmt.Println("Starting server on port 4000...")
	http.ListenAndServe(":4000", r)
}