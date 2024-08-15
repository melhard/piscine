package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		return
	}
	var result []string
	res := ""
	for _, w := range os.Args[1] {
		if w != ' ' {
			res += string(w)
		} else if w == ' ' && res != "" {
			result = append(result, res)
			res = ""
		}
	}
	if res != "" {
		result = append(result, res)
	}

	for i := 0; i < len(result); i++ {
		if i < len(result) {
			fmt.Print(result[i] + "\t")
		} else {
			fmt.Print(result[i])
		}
	}
	fmt.Println()
}
