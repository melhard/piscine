package piscine

func BasicJoin(elems []string) string {
	news := ""

	for _, s := range elems {
		news = news + s
	}
	return news
}
