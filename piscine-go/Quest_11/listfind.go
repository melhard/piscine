package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	node := l.Head
	for node != nil {
		if node == ref {
			comp(node, ref)
		}
		node = node.Next
	}
	return &ref
}
