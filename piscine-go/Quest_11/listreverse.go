package piscine

func ListReverse(l *List) {
	var yosf *NodeL
	jwhr := l.Head
	l.Tail = l.Head

	for jwhr != nil {
		next := jwhr.Next
		jwhr.Next = yosf
		yosf = jwhr
		jwhr = next
	}
	l.Head = yosf
}
