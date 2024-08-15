package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	elm := l
	index := 0
	for elm != nil {
		if index == pos {
			return elm
		}
		elm = elm.Next
		index++
	}
	return nil
}
