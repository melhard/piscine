package piscine

import (
	"fmt"
	"unicode"
)

func PrintMemory(arr [10]byte) {
	ps := "0123456789abcdef"
	res := ""
	cont := 1
	for _, ww := range arr {
		if cont == 4 {
			res += string(ps[ww/16]) + string(ps[ww%16])
			fmt.Println(res)
			res = ""
			cont = 1
		} else {
			res += string(ps[ww/16]) + string(ps[ww%16]) + " "
			cont++
		}
	}
	fmt.Println(res)

	for _, w := range arr {
		if unicode.IsGraphic(rune(w)) {
			fmt.Print(string(w))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}
