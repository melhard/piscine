package piscine

func ListSize(l *List) int {
	cont := 0
	elm := l.Head

	for elm != nil {
		cont++
		elm = elm.Next
	}

	return cont
}
