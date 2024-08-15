package piscine

func ToLower(s string) string {
	slice := []rune(s)
	for i, worrd := range s {
		if worrd >= 65 && worrd <= 90 {
			slice[i] = worrd - ('A' - 'a')
		} else if worrd == 33 || worrd == 63 {
			slice[i] = worrd
		}
	}
	return (string(slice))
}
