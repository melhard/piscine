package piscine

func ForEach(f func(int), a []int) {
	for _, worrd := range a {
		f(worrd)
	}
}
