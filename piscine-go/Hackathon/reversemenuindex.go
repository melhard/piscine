package piscine

func ReverseMenuIndex(menu []string) []string {
	result := make([]string, len(menu))
	cont := 0
	for i := len(menu) - 1; i >= 0; i-- {
		result[cont] = menu[i]
		cont++
	}
	return result
}
