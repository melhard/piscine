package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	te := food{}
	if order == "burger" {
		te.preptime = 15
	} else if order == "chips" {
		te.preptime = 10
	} else if order == "nuggets" {
		te.preptime = 12
	} else {
		return 404
	}
	return te.preptime
}
