package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := os.Getenv("PATH")
	paths := strings.Split(path, ";")
	for i := 0; i < len(paths); i++ {
        fmt.Println(paths[i])
    }
}