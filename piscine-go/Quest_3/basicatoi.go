package piscine

func BasicAtoi(s string) int {
	me := 0
	for _, worrd := range s {
		me = me*10 + int(worrd-'0')
	}
	return me
}
