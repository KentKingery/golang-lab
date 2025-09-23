package main

import (
	"time"
)

type person struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"price"`
}
