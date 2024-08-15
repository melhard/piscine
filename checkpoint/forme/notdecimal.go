package piscine

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}
	res := ""
	start := 0
	negative := false
	if dec[0] == '-' {
		start = 1
		negative = true
	}
	s := 0

	if isn(dec) {
		for i := start; i < len(dec[s:]); i++ {
			if dec[i] == '.' {
				continue
			} else {
				res += string(dec[i])
			}
		}
		for j, ss := range res {
			if ss > '0' && ss <= '9' {
				s = j
				break
			}
		}
		if negative {
			return "-" + res[s:] + "\n"
		} else {
			return res[s:] + "\n"
		}
	} else {
		return dec
	}
}

func isn(s string) bool {
	for i, ww := range s {
		if i == 0 && ww == '-' {
			continue
		}
		if ww <= '0' && ww >= '9' {
			return false
		}
	}
	return true
}
