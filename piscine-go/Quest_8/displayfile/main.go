package main

import (
	"fmt"
	"os"
)

func main() {
	nm := os.Args[1:]
	if len(nm) == 0 {
		fmt.Println("File name missing")
	} else if len(nm) == 1 {
		fmt.Println("Almost there!!")
	} else {
		fmt.Println("Too many arguments")
	}
}
