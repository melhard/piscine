package main

import (
	"github.com/01-edu/z01"
)

const str = "x = 42, y = 21\n"

func main() {
	for _, char := range str {
		z01.PrintRune(char)
	}
}
