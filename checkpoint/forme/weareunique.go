package piscine

func WeAreUnique(str1, str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}
	help := false
	help1 := false
	cont := 0
	for j := rune(0); j < rune(127); j++ {
		for _, ww := range str1 {
			if ww == rune(j) {
				help = true
			}
		}
		for _, www := range str2 {
			if www == rune(j) {
				help1 = true
			}
		}
		if help != help1 {
			cont++
		}
		help = false
		help1 = false

	}

	return cont
}
