package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 2 {
		return
	}
	f1 := os.Args[1]
	f2 := os.Args[2]
	help := make(map[rune]int)
	res := ""

	for _, ww := range f1 {
		if help[ww] > 0 {
			continue
		}
		for i := 0; i < len(f2); i++ {
			if byte(ww) == f2[i] {
				help[ww]++
				res += string(ww)
				break
			}
		}
	}
	fmt.Println(res)
}
