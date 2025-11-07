package piscine

func ListReverse(l *List) {
	Prev := &NodeL{}
	Curr := &NodeL{}
	Prev = nil
	Curr = l.Head
	for Curr != nil {
		Nex := Curr.Next
		Curr.Next = Prev
		Prev = Curr
		Curr = Nex
	}
	l.Tail = l.Head
	l.Head = Prev
}
