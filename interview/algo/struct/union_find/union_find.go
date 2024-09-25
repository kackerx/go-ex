package union_find

type UF interface {
	GetSize() int
	IsConnected(p, q int) bool
	UnionElements(p, q int)
}

type UnionFind2 struct {
	parent []int
}

func NewUnionFind2(size int) *UnionFind2 {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &UnionFind2{parent}
}

func (u *UnionFind2) GetSize() int {
	return len(u.parent)
}

// find o(h), 查找p所对应的集合编号(根节点代表集合编号)
func (u *UnionFind2) find(p int) int {
	for p != u.parent[p] {
		p = u.parent[p]
	}

	return p
}

func (u *UnionFind2) IsConnected(p, q int) bool {
	return u.find(p) == u.find(q)
}

// UnionElements p元素的根节点 --> q元素的根节点
func (u *UnionFind2) UnionElements(p, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot == qRoot {
		return
	}

	u.parent[pRoot] = qRoot
}
