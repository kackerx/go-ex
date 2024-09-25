package tree

import (
	"fmt"
	"testing"
)

var arr = []int{-2, 0, 3, -5, 2, -1}

func TestBST(t *testing.T) {
	tree := NewBST()

	tree.Add(5)
	tree.Add(3)
	tree.Add(1)
	tree.Add(4)
	tree.Add(7)

	tree.InOrder()
}

func MaxDepth(no *node) int {
	if no.left == nil && no.right == nil {
		return 1
	}

	return max(MaxDepth(no.left), MaxDepth(no.right)) + 1
}

func TestSegment(t *testing.T) {
	s := NewSegmentTree(arr, func(a, b int) int {
		return a + b
	})

	fmt.Println(s.Query(2, 5))
}

func TestTrie(t *testing.T) {
	tr := NewTrie()

	tr.Insert("hello")
	tr.Insert("hel")
	tr.Insert("world")
	tr.Insert("woad")

	// fmt.Println(tr.Contains("hello"))
	// fmt.Println(tr.Contains("helo"))
	// fmt.Println(tr.Contains("world"))
	// fmt.Println(tr.Contains("wo.d"))
	//
	// fmt.Println(tr.ContainsRe("hello"))
	// fmt.Println(tr.ContainsRe("helo"))
	// fmt.Println(tr.ContainsRe("world"))
	// fmt.Println(tr.ContainsRe("word"))
	//
	// fmt.Println(tr.Query("he"))
	//
	// fmt.Println(tr.StartWith("h"))
	// fmt.Println(tr.StartWith("he"))
	// fmt.Println(tr.StartWith("heo"))

	fmt.Println(tr.ContainsRe("wo.d"))
	fmt.Println(tr.ContainsRe("wod"))
}

func TestTriee(t *testing.T) {
	tr := NewTriee()
	tr.Insert("hello")
	tr.Insert("hel")
	tr.Insert("world")
	tr.Insert("hhe")

	fmt.Println(tr.Contains("hel"))
	fmt.Println(tr.HasPrefix("he"))
	fmt.Println(tr.PassCnt("h"))
}
