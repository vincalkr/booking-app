package main

import (
	"booking-app/helpers"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(helpers.SplitFields("Hello World"))
	fmt.Println(Book{Title: "Aladin", Author: "WaltDisney"})
	fmt.Println(Book{Title: "La sirenetta", Author: "WaltDisney"})
}
