package piscine

func ZipString(s string) string {
	help := make(map[rune]int)
	var res []rune
	for i, ww := range s {
		if i != len(s)-1 {
			if s[i] == s[i]+1 {
				help[ww]++
			} else {
				help[ww]++
				res = append(res, rune(help[ww]+'0'))
				res = append(res, ww)
				help[ww] = 0
			}
		} else {
			help[ww]++
			res = append(res, rune(help[ww]+'0'))
			res = append(res, ww)
		}
	}
	return string(res)
}
