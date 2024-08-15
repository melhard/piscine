package piscine

func Prime(nb int) bool {
	if nb <= 1 {
		return false
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
}

func FindNextPrime(nb int) int {
	if Prime(nb) {
		return nb
	} else {
		return FindNextPrime(nb + 1)
	}
}
