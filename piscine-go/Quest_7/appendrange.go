package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	var mynewslice []int
	for i := min; i < max; i++ {
		mynewslice = append(mynewslice, i)
	}
	return mynewslice
}
