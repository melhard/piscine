package piscine

func IsUpper(s string) bool {
	count := 0

	for _, worrd := range s {
		if worrd >= 'A' && worrd <= 'Z' {
			count++
		}
	}
	return count == len(s)
}
