package piscine

func BasicAtoi2(s string) int {
	me := 0
	for _, worrrd := range s {
		if worrrd > 47 && worrrd < 58 { //[48......57]
			me = me*10 + int(worrrd-'0')
		} else {
			return 0
		}
	}
	return me
}
