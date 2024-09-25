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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	q := &queue{}
	q.Put(root)

	for !q.IsEmpty() {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			no := q.Get()
			if i == sz-1 {
				res = append(res, no.Val)
			}

			if no.Left != nil {
				q.Put(no.Left)
			}

			if no.Right != nil {
				q.Put(no.Right)
			}
		}
	}

	return res
}

type queue []*TreeNode

func (q *queue) Get() *TreeNode {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *queue) Put(no *TreeNode) {
	*q = append(*q, no)
}

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Len() int {
	return len(*q)
}

// leetcode submit region end(Prohibit modification and deletion)
