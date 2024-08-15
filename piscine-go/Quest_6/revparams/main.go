package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nm := os.Args[1:]
	for i := len(nm) - 1; i >= 0; i-- {
		for _, v := range nm[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
