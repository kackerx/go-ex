package union_find

type UF struct {
	count  int   // 连通分量
	parent []int // 节点x的
	size   []int // 每个节点有几个子节点, 判断树的大小以便于, 小树 --> 大树更加平衡
}

func NewUF(n int) *UF {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &UF{count: n, parent: parent}
}

// find 找到某个元素的根节点
func (u *UF) find(x int) int {
	if x != u.parent[x] {
		x = u.parent[x]
	}

	return x
}

// Union 连通两个分量, p.root --> q.root
func (u *UF) Union(p, q int) {
	rootP := u.find(p)
	rootQ := u.find(q)

	if rootP == rootQ {
		return
	}

	if u.size[rootP] < u.size[rootQ] {
		u.parent[rootP] = rootQ
		u.size[rootQ] += u.size[rootP]
	} else {
		u.parent[rootQ] = rootP
		u.size[rootP] += u.size[rootQ]
	}
	u.count--
}

// Connected 连通一定有相同的根节点
func (u *UF) Connected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UF) Count() int {
	return u.count
}
