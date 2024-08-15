package piscine

func ListLast(l *List) interface{} {
	for l.Tail == nil {
		return nil
	}
	return l.Tail.Data
}
