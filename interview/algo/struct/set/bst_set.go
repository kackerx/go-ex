package set

import "go-bobo/interview/algo/struct/tree"

type BSTSet struct {
	bst *tree.BST
}

func (b *BSTSet) Add(e int) {
	b.bst.Add(e)
}

func (b *BSTSet) Remove(e int) {
	b.bst.Remove(e)
}

func (b *BSTSet) Contains(e int) {
	b.bst.Contains(e)
}

func (b *BSTSet) Size() int {
	return b.bst.Size()
}

func (b *BSTSet) IsEmpty() bool {
	return b.bst.IsEmpty()
}
