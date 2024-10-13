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
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}

	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, m)
}

func build(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, m map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	rootVal := preorder[preStart]
	rootIdx := m[rootVal]
	leftSize := rootIdx - inStart

	root := &TreeNode{
		Val:   rootVal,
		Left:  build(preorder, preStart+1, preStart+leftSize, inorder, inStart, rootIdx-1, m),
		Right: build(preorder, preStart+leftSize+1, preEnd, inorder, rootIdx+1, inEnd, m),
	}

	return root
}

// leetcode submit region end(Prohibit modification and deletion)
