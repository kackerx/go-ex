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
	sep  string
	null string
	res  []string
}

func Constructor() Codec {
	return Codec{",", "#", make([]string, 0)}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.traverse(root)
	res := strings.Join(this.res, this.sep)
	return res
}

func (this *Codec) traverse(no *TreeNode) {
	if no == nil {
		this.res = append(this.res, this.null)
		return
	}

	this.res = append(this.res, strconv.Itoa(no.Val))
	this.traverse(no.Left)
	this.traverse(no.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, this.sep)
	return this.Deserializes(&nodes)
}

func (this *Codec) Deserializes(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}

	first := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if first == this.null {
		return nil
	}

	val, _ := strconv.Atoi(first)
	no := &TreeNode{Val: val}

	no.Left = this.Deserializes(nodes)
	no.Right = this.Deserializes(nodes)

	return no
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// leetcode submit region end(Prohibit modification and deletion)
