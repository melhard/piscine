package main

import (
	"fmt"
	"os"
)

func main() {
	nm := os.Args[1:]
	for _, worrd := range nm {
		if worrd == "01" || worrd == "galaxy" || worrd == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
