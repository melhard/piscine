package piscine

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	for b != 0 {
		r := a % b
		a = b
		b = r
	}
	return a
}
