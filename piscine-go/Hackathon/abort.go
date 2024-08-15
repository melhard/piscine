package piscine

func Abort(a, b, c, d, e int) int {
	result := []int{a, b, c, d, e}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] < result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result[2]
}
