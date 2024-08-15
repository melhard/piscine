package piscine

func Itoaa(n int) string {
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	res := ""
	result := ""
	for n > 0 {
		help := n % 10
		res = res + string(rune(help+'0'))
		n = n / 10
	}

	for i := len(res) - 1; i >= 0; i-- {
		result += string(res[i])
	}

	if negative {
		return "-" + result
	}
	return result
}
