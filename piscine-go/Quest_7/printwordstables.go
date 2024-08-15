package piscine

import "github.com/01-edu/z01"

func ConcatParamss(args []string) string {
	var res string

	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			res += args[i]
		} else {
			res += args[i] + "\n"
		}
	}
	return res
}

func PrintWordsTables(a []string) {
	mynm := (ConcatParamss(a))
	for _, worrd := range mynm {
		z01.PrintRune(worrd)
	}
	z01.PrintRune('\n')
}
