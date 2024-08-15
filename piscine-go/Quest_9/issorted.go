package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var result int
	for i := range a {
		if i < len(a)-3 {
			result = f(a[i], a[i+1])
			if (result > 0 && f(a[i+1], a[i+2]) < 0) || (result < 0 && f(a[i+1], a[i+2]) > 0) {
				return false
			}
		}
	}
	return true
}
