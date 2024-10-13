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

	var res []int
	q := &Queue{}
	q.Enqueue(root)

	for !q.IsEmpty() {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			no := q.Dequeue()
			if i == sz-1 {
				res = append(res, no.Val)
			}

			if no.Left != nil {
				q.Enqueue(no.Left)
			}

			if no.Right != nil {
				q.Enqueue(no.Right)
			}
		}
	}

	return res
}

type Queue []*TreeNode

func (q *Queue) Dequeue() *TreeNode {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func (q *Queue) Enqueue(node *TreeNode) {
	*q = append(*q, node)
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

// leetcode submit region end(Prohibit modification and deletion)
