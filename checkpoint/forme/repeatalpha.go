package piscine

func RepeatAlpha(s string) string {
	res := ""
	for _, ww := range s {
		if ww >= 'a' && ww <= 'z' {
			for i := 'a'; i <= ww; i++ {
				res += string(ww)
			}
		} else if ww >= 'A' && ww <= 'Z' {
			for i := 'A'; i <= ww; i++ {
				res += string(ww)
			}
		} else {
			res += string(ww)
		}
	}
	return res
}
