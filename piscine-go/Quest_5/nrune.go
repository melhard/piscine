package piscine

func NRune(s string, n int) rune {
	slice := []rune(s)
	if n < 0 || n > len(s) {
		return 0
	} else {
		for i, worrd := range slice {
			if i == n-1 {
				return worrd
			}
		}
		return 0
	}
}
