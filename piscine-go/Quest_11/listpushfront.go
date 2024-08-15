package piscine

func ListPushFront(l *List, data interface{}) {
	newnode := &NodeL{
		Data: data,
	}
	new := l.Head
	if l.Tail == nil {
		l.Head = newnode
		l.Tail = newnode
	} else {
		l.Head = newnode
		l.Head.Next = new
	}
}
