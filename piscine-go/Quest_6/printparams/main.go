package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nm := os.Args[1:]
	for _, wo := range nm {
		for _, v := range wo {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
