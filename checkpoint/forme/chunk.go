package piscine

import "fmt"

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		fmt.Println(slice)
		return
	}
	var help [][]int
	i := 0
	for i < len(slice) {
		if i+size < len(slice) {
			boob := slice[i : i+size]
			help = append(help, boob)
			i += size
		} else {
			boob := slice[i:]
			help = append(help, boob)
			break
		}
	}
	fmt.Println(help)
}
