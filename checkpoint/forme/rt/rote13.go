package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	new := os.Args[1]
	for _, ww := range new {
		res := Rote(ww)
		z01.PrintRune(res)
	}
	z01.PrintRune('\n')
}

func Rote(ww rune) rune {
	switch {
	case ww >= 'a' && ww <= 'n' ||  ww >= 'A' && ww <= 'N' :
		return ww + 13
	case ww >= 'm' && ww <= 'z' || ww >= 'M' && ww <= 'Z':
		return ww - 13
	default:
		return ww
	}
}
