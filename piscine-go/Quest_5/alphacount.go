package piscine

func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		if (char <= 'z' && char >= 'a') || (char <= 'Z' && char >= 'A') {
			count++
		}
	}
	return count
}
