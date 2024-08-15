package piscine

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	var mynewstring string

	if from <= to {
		for i := from; i <= to; i++ {
			mynewstring += string(rune('0'+i/10)) + string(rune('0'+i%10))
			if i < to {
				mynewstring += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			mynewstring += string(rune('0'+i/10)) + string(rune('0'+i%10))
			if i > to {
				mynewstring += ", "
			}
		}
	}
	mynewstring += "\n"
	return mynewstring
}
