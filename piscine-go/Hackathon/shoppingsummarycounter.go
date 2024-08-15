package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	list := make(map[string]int)
	var keys string
	for _, e := range str {
		if e == ' ' {
			list[keys] += 1
			keys = ""
		} else {
			keys += string(e)
		}
	}
	list[keys] += 1
	return list
}
