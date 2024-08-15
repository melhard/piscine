package piscine

func IsNumeric(s string) bool {
	var count int
	for _, worrd := range s {
		if (worrd >= '0') && (worrd <= '9') {
			count++
		}
	}
	return count == len(s)
}
