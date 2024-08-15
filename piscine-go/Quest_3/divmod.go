package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b // 13 / 2 = 6
	*mod = a % b // 1
}
