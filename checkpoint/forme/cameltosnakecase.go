package piscine

func IsUpper(ww rune) bool {
	if ww >= 'A' && ww <= 'Z' {
		return true
	}
	return false
}

func CamelToSnakeCase(s string) string {
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return s
	}
	res := ""
	for i := 0; i < len(s); i++ {
		if i != 0 && IsUpper(rune(s[i])) && !IsUpper(rune(s[i+1])) && i < len(s) {
			res += "_"
		}
		res += string(s[i])
	}
	return res
}
