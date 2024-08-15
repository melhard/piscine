package piscine

func IsPrintable(s string) bool {
	var count int
	for _, worrd := range s {
		if (worrd >= 32) && (worrd <= 126) {
			count++
		}
	}
	return count == len(s)
}
