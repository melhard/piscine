package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	new := make([]int, max-min)
	for i := min; i < max; i++ {
		new[i-min] = i
	}
	return new
}
