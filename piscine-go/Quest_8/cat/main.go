package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	leng := len(os.Args)
	if leng == 1 {
		res := make([]byte, 1000)
		n, _ := os.Stdin.Read(res)
		if n > 0 {
			res = res[:n-1]
			for string(res) == "Hello" {
				printString(string(res) + "\n")
			}
			printString(string(res) + "\n")
		}
	} else {
		for _, ww := range os.Args[1:] {
			file, err := os.Open(ww)
			if err != nil {
				printString("ERROR: ")
				printString(err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			}
			stat, _ := file.Stat()
			arr := make([]byte, stat.Size())
			file.Read(arr)
			printString(string(arr))
			file.Close()
		}
	}
}
