package heap

type Node struct {
	val  int
	next *Node
}

type pp []*Node

func (p *pp) Len() int {
	return len(*p)
}

func (p *pp) Less(i, j int) bool {
	return (*p)[i].val < (*p)[j].val
}

func (p *pp) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *pp) Push(x any) {
	*p = append(*p, x.(*Node))
}

func (p *pp) Pop() any {
	var v *Node

	v = (*p)[len(*p)-1]
	*p = (*p)[:p.Len()-1]
	return v
}
