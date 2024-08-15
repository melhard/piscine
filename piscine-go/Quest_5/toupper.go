package piscine

func ToUpper(s string) string {
	slice := []rune(s)
	for i, worrd := range s {
		if worrd >= 'a' && worrd <= 'z' {
			slice[i] = worrd - ('a' - 'A')
		}
	}
	return (string(slice))
}
