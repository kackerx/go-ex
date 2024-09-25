package main

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	traverse(root.Left, root.Right)

	return root
}

func traverse(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	node1.Next = node2
	traverse(node1.Left, node1.Right)
	traverse(node2.Left, node2.Right)
	traverse(node1.Right, node2.Left)
}

// leetcode submit region end(Prohibit modification and deletion)
