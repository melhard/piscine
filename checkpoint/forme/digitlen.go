package piscine

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 0 {
		n = -n
	}
	var conut int
	for n > 0 {
		n /= base
		conut++
	}
	return conut
}
