package piscine

func IsAlpha(s string) bool {
	var count int
	for _, worrd := range s {
		if (worrd >= 'A' && worrd <= 'Z') || (worrd >= 'a' && worrd <= 'z') || (worrd >= '0' && worrd <= '9') {
			count++
		}
	}
	return count == len(s)
}
