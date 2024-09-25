package tree

import (
	"fmt"
	"strings"

	"go-bobo/interview/algo/struct/queue"
	"go-bobo/interview/algo/struct/stack"
)

type node struct {
	data        int
	left, right *node
}

func newNode(data int, left *node, right *node) *node {
	return &node{data: data, left: left, right: right}
}

type BST struct {
	root *node
	size int
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) Size() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

func (b *BST) Add(e int) {
	b.root = b.add(b.root, e)
}

// add 返回插入新节点后的二分搜索树的根
func (b *BST) add(n *node, e int) *node {
	if n == nil {
		b.size++
		return newNode(e, nil, nil)
	}

	if e < n.data {
		n.left = b.add(n.left, e)
	} else if e > n.data {
		n.right = b.add(n.right, e)
	}

	return n
}

func (b *BST) Contains(e int) bool {
	return b.contains(b.root, e)
}

// contains 以n为根节点的树中是否包含e
func (b *BST) contains(n *node, e int) bool {
	if n == nil {
		return false
	}

	if n.data == e {
		return true
	} else if n.data < e {
		return b.contains(n.right, e)
	} else {
		return b.contains(n.left, e)
	}
}

func (b *BST) PreOrd() {
	s := stack.NewArrayStock[*node](10)

	s.Push(b.root)

	for !s.IsEmpty() {
		no := s.Pop()
		fmt.Println(no.data)

		if no.right != nil {
			s.Push(no.right)
		}

		if no.left != nil {
			s.Push(no.left)
		}

	}
}

func (b *BST) LevelOrder() {
	q := queue.NewArrayQueue[*node](10)
	q.Enqueue(b.root)
	var depth int

	for !q.IsEmpty() {
		fmt.Println(depth)

		no := q.Dequeue()
		fmt.Println(no.data)

		if no.left != nil {
			q.Enqueue(no.left)
		}

		if no.right != nil {
			q.Enqueue(no.right)
		}

		depth++
	}
}

func (b *BST) PreOrder() {
	b.preOrder(b.root)
}

func (b *BST) preOrder(n *node) {
	if n == nil {
		return
	}

	// 做访问该节点时该做的事情
	fmt.Println(n.data)

	b.preOrder(n.left)
	b.preOrder(n.right)
}

// InOrder 结果是升序有序排列
func (b *BST) InOrder() {
	b.inOrder(b.root)
}

func (b *BST) inOrder(n *node) {
	if n == nil {
		return
	}

	b.inOrder(n.left)

	// 做访问该节点时该做的事情
	fmt.Println(n.data)

	b.inOrder(n.right)
}

func (b *BST) PostOrder() {
	b.postOrder(b.root)
}

func (b *BST) postOrder(n *node) {
	if n == nil {
		return
	}

	b.postOrder(n.left)
	b.postOrder(n.right)

	// 做访问该节点时该做的事情
	fmt.Println(n.data)
}

func (b *BST) Maximum() int {
	return b.maximum(b.root).data
}

func (b *BST) maximum(no *node) *node {
	if no.right == nil {
		return no
	}

	return b.maximum(no.right)
}

func (b *BST) Minimum() int {
	return b.minimum(b.root).data
}

func (b *BST) minimum(no *node) *node {
	if no.left == nil {
		return no
	}

	return b.minimum(no.left)
}

func (b *BST) RemoveMax() int {
	maximum := b.Maximum()

	b.root = b.removeMax(b.root)
	return maximum
}

func (b *BST) removeMax(no *node) *node {
	if no.right == nil {
		return no.left
	}

	no.right = b.removeMax(no.right)
	return no
}

func (b *BST) RemoveMin() int {
	minimum := b.Minimum()

	b.root = b.removeMin(b.root)
	return minimum
}

// removeMin 删除no节点的最小节点, 返回删除后新的bst的根
func (b *BST) removeMin(no *node) *node {
	if no.left == nil {
		return no.right
	}

	no.left = b.removeMin(no.left)
	return no
}

func (b *BST) Remove(e int) {
	b.root = b.remove(b.root, e)
}

// remove 删除node为根的元素为e的节点, 返回删除后的新的bst的根
func (b *BST) remove(no *node, e int) *node {
	if no == nil {
		return no
	}

	if e < no.data {
		no.left = b.remove(no.left, e)
		return no
	} else if e > no.data {
		no.right = b.remove(no.right, e)
		return no
	}

	if no.left == nil {
		return no.right
	}

	if no.right == nil {
		return no.left
	}

	// 找到右子树的最小值
	successor := b.minimum(no.right)
	// 继任节点的右子树 --> 删除右子树最小值后的bst
	successor.right = b.removeMin(no.right)
	// 继任节点的左子树 --> 当前节点左子树
	successor.left = no.left
	// 当前节点置空
	no.left, no.right = nil, nil
	return successor
}

func (b *BST) String() string {
	builder := &strings.Builder{}
	b.string(b.root, 0, builder)

	return builder.String()
}

func (b *BST) string(n *node, depth int, res *strings.Builder) {
	if n == nil {
		res.WriteString("nil\n")
		return
	}

	res.WriteString(fmt.Sprintf("%s%d\n", b.genDepthString(depth), n.data))
	b.string(n.left, depth+1, res)
	b.string(n.right, depth+1, res)
}

func (b *BST) genDepthString(depth int) string {
	var res string
	for i := 0; i < depth; i++ {
		res += "--"
	}

	return res
}
