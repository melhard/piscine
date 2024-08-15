package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println()
		return
	}
	res := ""
	var result []string
	for _, ww := range os.Args[1] {
		if ww != ' ' {
			res += string(ww)
		} else if ww == ' ' && res != "" {
			result = append(result, res)
			res = ""
		}
	}
	if res != "" {
		result = append(result, res)
	}

	for i := 0; i < len(result); i++ {
		if i < len(result) {
			fmt.Print(result[i] + " ")
		} else {
			fmt.Print(result[i])
		}
	}
	fmt.Println()
}
