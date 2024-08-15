package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var slice []rune

	for n > 0 {
		mo := n % 10
		slice = append(slice, rune(mo+'0'))
		n = n / 10
	}
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	for _, char := range slice {
		z01.PrintRune(char)
	}
}
