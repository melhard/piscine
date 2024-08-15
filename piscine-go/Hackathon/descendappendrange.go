package piscine

func DescendAppendRange(max, min int) []int {
	if min >= max {
		return []int{}
	}

	var mynewslice []int = []int{}
	for i := max; i > min; i-- {
		mynewslice = append(mynewslice, i)
	}
	return mynewslice
}
