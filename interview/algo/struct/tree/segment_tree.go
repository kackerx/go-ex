package tree

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type SegmentTree struct {
	data   []int
	tree   []int
	merger func(a, b int) int
}

func NewSegmentTree(arr []int, merge func(a, b int) int) *SegmentTree {
	s := &SegmentTree{data: slices.Clone(arr), tree: make([]int, 4*len(arr)), merger: merge}
	s.build(0, 0, len(s.data)-1)

	return s
}

func (s *SegmentTree) Size() int {
	return len(s.data)
}

func (s *SegmentTree) leftChild(index int) int {
	return 2*index + 1
}

func (s *SegmentTree) rightChild(index int) int {
	return 2*index + 2
}

// build 在treeIndex位置, 构造表示区间[l...r]的线段树
func (s *SegmentTree) build(treeIndex, l, r int) {
	if l == r {
		s.tree[treeIndex] = s.data[l]
		return
	}

	leftIndex := s.leftChild(treeIndex)
	rightIndex := s.rightChild(treeIndex)

	mid := l + (r-l)/2
	s.build(leftIndex, l, mid)
	s.build(rightIndex, mid+1, r)

	s.tree[treeIndex] = s.merger(s.tree[leftIndex], s.tree[rightIndex])
}

func (s *SegmentTree) Query(queryL, queryR int) int {
	return s.query(0, 0, len(s.data)-1, queryL, queryR)
}

// Query 以treeIndex为根的线段树中[l, r]的范围里查找queryL, queryR的结果
func (s *SegmentTree) query(treeIndex, l, r, queryL, queryR int) int {
	if queryL == l && queryR == r {
		return s.tree[treeIndex]
	}

	mid := l + (r-l)/2
	leftIndex := s.leftChild(treeIndex)
	rightIndex := s.rightChild(treeIndex)
	if queryL >= mid+1 {
		return s.query(rightIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid {
		return s.query(leftIndex, l, mid, queryL, queryR)
	}

	leftRes := s.query(leftIndex, l, mid, queryL, mid)
	rightRes := s.query(rightIndex, mid+1, r, mid+1, queryR)

	return s.merger(leftRes, rightRes)
}

func (s *SegmentTree) String() string {
	return fmt.Sprintf("%v", s.tree)
}
