package piscine

func TrimAtoi(s string) int {
	n := 0
	sign := 1
	for _, char := range s {
		if IsNumeric(string(char)) {
			n = n*10 + int(char-48)
		} else if char == '-' && n == 0 {
			sign = sign * (-1)
		}
	}
	return n * sign
}
