package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 4 {
		return
	}

	f1 := os.Args[1]
	f2 := os.Args[2]

	i := 0
	j := 0

	for i < len(f1) && j < len(f2) {
		if f1[i] == f2[j] {
			i++
		} else {
			j++
		}
	}
	if i != len(f1)-1 {
		fmt.Println(f1)
	}
}
