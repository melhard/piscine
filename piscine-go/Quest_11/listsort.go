package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	help := true
	for help {
		yosf := l
		help = false
		for yosf != nil && yosf.Next != nil {
			if yosf.Data > yosf.Next.Data {
				yosf.Data, yosf.Next.Data = yosf.Next.Data, yosf.Data
				help = true
			}
			yosf = yosf.Next
		}
	}
	return l
}
