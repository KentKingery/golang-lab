package main

import (
	"fmt"

	"acmecode/basic/models"
)

func main() {
	fmt.Println("Hello, World!")
	c := models.Customer{ID: 1, Name: "John Doe", Age: 30, Email: "M0z1c@example.com"}
	fmt.Println(c.Name)
}
