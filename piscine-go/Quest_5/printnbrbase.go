package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if nbr == -9223372036854775808 {
		PrintNbrBase(-922337203685477580, base)
		z01.PrintRune('8')
		return
	}
	res := ""
	isnegative := false
	if isvalide(base) {
		if nbr < 0 {
			nbr = -nbr
			isnegative = true
		}

		for nbr > 0 {
			res = string(base[nbr%len(base)]) + res
			nbr = nbr / len(base)
		}
		if isnegative {
			res = "-" + res
		}
		for _, char := range res {
			z01.PrintRune(char)
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}

func isvalide(base string) bool {
	if len(base) < 2 {
		return false
	}
	cont := 0
	for i := rune(0); i <= rune(127); i++ {
		for _, ww := range base {
			if ww == i {
				cont++
			}
			if ww == '-' || ww == '+' {
				return false
			}
		}
		if cont > 1 {
			return false
		}
		cont = 0
	}
	return true
}
