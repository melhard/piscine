package piscine

func ThirdTimeIsACharm(str string) string {
	if len(str) <= 2 {
		return "\n"
	}
	var res string
	for i := 2; i < len(str); i += 3 {
		res += string(str[i])
	}
	return res + "\n"
}
