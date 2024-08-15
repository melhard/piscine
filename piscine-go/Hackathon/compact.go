package piscine

func Compact(ptr *[]string) int {
	var result []string
	cont := 0
	for _, woorrd := range *ptr {
		if woorrd != "" {
			result = append(result, woorrd)
			cont++
		}
	}
	*ptr = result
	return cont
}
