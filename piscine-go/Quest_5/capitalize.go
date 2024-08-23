package piscine

func isPrint(r rune) bool {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	mySlice := []rune(s)

	next := true

	for i := 0; i < len(s); i++ {
		if isPrint(mySlice[i]) && next {
			if mySlice[i] >= 'a' && mySlice[i] <= 'z' {
				mySlice[i] = mySlice[i] - ('a' - 'A')
			}
			next = false
		} else if mySlice[i] >= 'A' && mySlice[i] <= 'Z' {
			mySlice[i] = mySlice[i] + ('a' - 'A')
		} else if !isPrint(mySlice[i]) {
			next = true
		}
	}
	return string(mySlice)
}


/*package main

import "fmt"

func main() {
	fmt.Println(Capitalize("hello! How are you? how+are+things+4you?"))
}
func isprint(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return false
	}
	return true
}
func Capitalize(s string) string {
	ss := []rune(s)
	if ss[0] >= 'a' && ss[0] <= 'z' {
		ss[0] = ss[0] - ('a' - 'A')
	}
	for i := 0; i < len(s); i++ {
		if ss[i] >= 'a' && ss[i] <= 'z' && isprint(ss[i-1]) {
			ss[i] = ss[i] - ('a' - 'A')
		}
	}
	return string(ss)
}*/
