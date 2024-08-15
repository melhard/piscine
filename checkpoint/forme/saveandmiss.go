package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}
	res := ""
	cont := 0

	for i := 0; i < len(arg); i++ {
		if cont < num {
			res += string(arg[i])
			cont++
		} else if cont == num {
			i += num - 1
			cont = 0
		}
	}
	return res
}
