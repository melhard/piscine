package piscine

func SplitWhiteSpaces(s string) []string {
	var myres []string
	var ww string // to split
	for _, woord := range s {
		if woord == ' ' || woord == '\t' || woord == '\n' {
			if ww != "" {
				myres = append(myres, ww)
				ww = ""
			}
		} else {
			ww += string(woord)
		}
	}
	if ww != "" { // to make ww empty agian
		myres = append(myres, ww)
	}
	return myres
}
