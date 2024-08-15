package piscine

func Atoi(s string) int {
	negative := false
	start := 0
	result := 0
	new := []rune(s)

	if len(s) == 0 {
		return 0
	}

	if new[0] == '-' {
		negative = true
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(new); i++ {
		if new[i] < '0' || new[i] > '9' {
			return 0
		}
		result = result*10 + int(new[i]-'0')
	}
	if negative {
		return -result
	}
	return result
}
