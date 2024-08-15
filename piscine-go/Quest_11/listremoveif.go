package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	node := l.Head
	for node != nil && node.Next != nil {
		if node.Next.Data == data_ref {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
}
