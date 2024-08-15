package piscine

func AtoiBase(s string, base string) int {
	if isvalid1(base) {
		res := 0
		s = StrRev1(s)
		for j, nb := range s {
			for i, ba := range base {
				if ba == nb {
					res += i * IterativePower1(len(base), j)
				}
			}
		}
		return res
	}
	return 0
}

func StrRev1(s string) string {
	ss := []byte(s)
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	result := (string(ss))
	return result
}

func isvalid1(base string) bool {
	if len(base) < 2 {
		return false
	}
	c := 0
	for i := rune(0); i <= rune(127); i++ {
		for _, char := range base {
			if char == i {
				c++
			}
		}
		if c > 1 {
			return false
		}
		c = 0
	}
	for _, char := range base {
		if char == '-' || char == '+' {
			return false
		}
	}
	return true
}

func IterativePower1(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		result := nb
		for i := 1; i < power; i++ {
			result = result * nb
		}
		return result
	}
}
