package piscine

func CamelToSnakeCasee(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' && s[i+1] >= 'A' && s[i+1] <= 'Z' {
			return s
		} else if s[i] >= '0' && s[i] <= '9' {
			return s
		}
	}
	res := ""
	for j, ww := range s {
		if ww >= 'A' && ww <= 'Z' && j != 0 {
			res += "_" + string(ww)
		} else {
			res += string(ww)
		}
	}
	return res
}
