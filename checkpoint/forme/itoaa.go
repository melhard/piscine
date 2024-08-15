package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	negative := n < 0
	if negative {
		n = -n
	}
	res := ""
	help := 0

	for n != 0 {
		help = n % 10
		res = string(rune(help+'0')) + res
		n = n / 10
	}

	if negative {
		res = "-" + res
	}

	return string(res)
}
