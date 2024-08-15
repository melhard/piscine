package piscine

func CanJump(s []uint) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); {
		if i == len(s)-1 || len(s) == 1 {
			return true
		}
		if s[i] == 0 {
			return false
		}
		i += int(s[i])
	}
	return false
}
