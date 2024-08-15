package piscine

func IsLower(s string) bool {
	count := 0

	for _, worrd := range s {
		if worrd >= 'a' && worrd <= 'z' {
			count++
		}
	}
	return count == len(s)
}
