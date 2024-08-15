package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil && l2.Head == nil {
		return
	} else if l1.Head != nil && l2.Head == nil {
		return
	} else if l1.Head == nil && l2.Head != nil {
		l1.Head = l2.Head
	} else if l1.Head != nil && l2.Head != nil {
		l1.Tail = l2.Head.Next
	}
}
