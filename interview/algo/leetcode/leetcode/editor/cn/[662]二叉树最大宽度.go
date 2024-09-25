package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Pair struct {
	node *TreeNode
	id   int
}

type queue []*Pair

func (q *queue) Get() *Pair {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *queue) Put(no *Pair) {
	*q = append(*q, no)
}

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Len() int {
	return len(*q)
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return -1
	}

	var maxWidth, left, right int
	q := &queue{}
	q.Put(&Pair{root, 1})

	for !q.IsEmpty() {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			no := q.Get()
			if i == 0 {
				left = no.id
			}
			if i == sz-1 {
				right = no.id
			}

			if no.node.Left != nil {
				q.Put(&Pair{no.node.Left, 2 * (no.id)})
			}

			if no.node.Right != nil {
				q.Put(&Pair{no.node.Right, 2*(no.id) + 1})
			}
		}
		maxWidth = max(maxWidth, right-left+1)
	}

	return maxWidth
}

// leetcode submit region end(Prohibit modification and deletion)
