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
func widthOfBinaryTree(root *TreeNode) int {
	q := &Queue{}
	maxWidth := 0
	q.Enqueue(&Pair{
		node: root,
		id:   1,
	})

	for !q.IsEmpty() {
		sz := q.Len()
		var start, end int
		for i := 0; i < sz; i++ {
			no := q.Dequeue()
			if i == 0 {
				start = no.id
			}
			if i == sz-1 {
				end = no.id
			}

			if no.node.Left != nil {
				q.Enqueue(&Pair{
					node: no.node.Left,
					id:   no.id * 2,
				})
			}
			if no.node.Right != nil {
				q.Enqueue(&Pair{
					node: no.node.Right,
					id:   no.id*2 + 1,
				})
			}
		}

		maxWidth = max(maxWidth, end-start+1)
	}

	return maxWidth
}

type Pair struct {
	node *TreeNode
	id   int
}

type Queue []*Pair

func (q *Queue) Dequeue() *Pair {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func (q *Queue) Enqueue(node *Pair) {
	*q = append(*q, node)
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

// leetcode submit region end(Prohibit modification and deletion)
