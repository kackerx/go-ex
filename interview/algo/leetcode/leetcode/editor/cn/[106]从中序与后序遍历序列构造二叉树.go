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
func buildTree(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int, 0)
	for i, v := range inorder {
		m[v] = i
	}

	return build(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1, m)
}

func build(inorder, postorder []int, inStart, inEnd, postStart, postEnd int, m map[int]int) *TreeNode {
	if inStart > inEnd {
		return nil
	}

	rootVal := postorder[postEnd]
	rootIdx := m[rootVal]
	node := &TreeNode{Val: rootVal}
	leftSize := rootIdx - inStart

	node.Left = build(inorder, postorder, inStart, rootIdx-1, postStart, postStart+leftSize-1, m)
	node.Right = build(inorder, postorder, rootIdx+1, inEnd, postStart+leftSize, postEnd-1, m)

	return node
}

// leetcode submit region end(Prohibit modification and deletion)
