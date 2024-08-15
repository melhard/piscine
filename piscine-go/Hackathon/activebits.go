package piscine

func ActiveBits(n int) int {
	var count int
	for n > 0 {
		mod := n % 2
		if mod != 0 {
			count++
		}
		n = n / 2
	}
	return count
}
