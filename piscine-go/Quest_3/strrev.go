package piscine

func StrRev(s string) string {
	slice := []byte(s)
	var rev string
	for i := len(slice) - 1; i >= 0; i-- {
		rev = rev + string(slice[i])
	}
	return rev
}
