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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := &Queue{}
	var res [][]int

	q.Enqueue(root)
	for !q.IsEmpty() {
		sz := q.Len()
		var levelRet []int

		for i := 0; i < sz; i++ {
			no := q.Dequeue()
			levelRet = append(levelRet, no.Val)

			if no.Left != nil {
				q.Enqueue(no.Left)
			}

			if no.Right != nil {
				q.Enqueue(no.Right)
			}
		}

		res = append(res, levelRet)
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
