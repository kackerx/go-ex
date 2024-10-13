package main

import (
	"strconv"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	sep, null string
}

func Constructor() Codec {
	return Codec{
		sep:  ",",
		null: "#",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	res := strings.Builder{}

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			res.WriteString(this.null)
			res.WriteString(this.sep)
			return
		}

		res.WriteString(strconv.Itoa(node.Val))
		res.WriteString(this.sep)
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)

	return res.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {return nil}
	nodeList := strings.Split(data[:len(data)-1], this.sep)
	return this._deserialize(&nodeList)
}

func (this *Codec) _deserialize(nodeList *[]string) *TreeNode {
	if len(*nodeList) == 0 {
		return nil
	}

	first := (*nodeList)[0]
	*nodeList = (*nodeList)[1:]
	if first == this.null {
		return nil
	}

	v, _ := strconv.Atoi(first)
	root := &TreeNode{Val: v}

	root.Left = this._deserialize(nodeList)
	root.Right = this._deserialize(nodeList)

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// leetcode submit region end(Prohibit modification and deletion)
