package piscine

func Unmatch(a []int) int {
	res := 0
	for _, ww := range a {
		for _, www := range a {
			if ww == www {
				res++
			}
		}
		if res%2 != 0 {
			return ww
		}
	}
	return -1
}
