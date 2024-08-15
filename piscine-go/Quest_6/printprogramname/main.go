package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nm := os.Args[0]
	for _, wo := range nm {
		if wo == '.' || wo == '/' {
			continue
		} else {
			z01.PrintRune(wo)
		}
	}
	z01.PrintRune('\n')
}
