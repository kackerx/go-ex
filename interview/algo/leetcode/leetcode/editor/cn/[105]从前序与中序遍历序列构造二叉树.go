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
func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int, 0)
	for i, v := range inorder {
		m[v] = i
	}

	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1, m)
}

func build(preorder []int, inorder []int, preStart int, preEnd int, inStart int, inEnd int, m map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	node := &TreeNode{Val: preorder[preStart]}

	nodeIndex := m[preorder[preStart]]
	leftSize := nodeIndex - inStart

	node.Left = build(preorder, inorder, preStart+1, preStart+leftSize, inStart, nodeIndex-1, m)
	node.Right = build(preorder, inorder, preStart+leftSize+1, preEnd, nodeIndex+1, inEnd, m)

	return node
}

// leetcode submit region end(Prohibit modification and deletion)
