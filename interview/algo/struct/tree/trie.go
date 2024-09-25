package tree

import "golang.org/x/exp/maps"

type Node struct {
	isWord bool
	next   map[byte]*Node
}

func NewNode(isWord bool) *Node {
	return &Node{isWord: isWord, next: make(map[byte]*Node)}
}

type Trie struct {
	root *Node
	size int
}

func NewTrie() *Trie {
	return &Trie{root: NewNode(false)}
}

// Size 单词数量
func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		if _, ok := cur.next[word[i]]; !ok {
			cur.next[word[i]] = NewNode(false)
		}
		cur = cur.next[word[i]]
	}

	// 此时cur指向单词最后一节点
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
}

func (t *Trie) Contains(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		if _, ok := cur.next[word[i]]; !ok {
			return false
		}
		cur = cur.next[word[i]]
	}

	return cur.isWord
}

func (t *Trie) ContainsRe(word string) bool {
	return t.containsRe(t.root, word, 0)
}

// containsRe 节点no为根节点的树中, 是否包含字符串word, 从索引index开始的部分, 支持正则.
func (t *Trie) containsRe(no *Node, word string, index int) bool {
	if len(word) == index {
		return no.isWord
	}

	if word[index] != '.' {
		next, ok := no.next[word[index]]
		if !ok {
			return false
		} else {
			return t.containsRe(next, word, index+1)
		}
	} else {
		for _, nextChar := range maps.Keys(no.next) {
			if t.containsRe(no.next[nextChar], word, index+1) {
				return true
			}
		}
		return false
	}
}

func (t *Trie) Query(prefix string) []string {
	var res []string
	cur := t.root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if _, ok := cur.next[char]; !ok {
			return res
		}

		cur = cur.next[char]
	}

	t.queryRe(cur, prefix, &res)

	return res
}

func (t *Trie) queryRe(no *Node, curPrefix string, res *[]string) {
	if no.isWord {
		*res = append(*res, curPrefix)
	}

	for c, nextNode := range no.next {
		t.queryRe(nextNode, curPrefix+string(c), res)
	}
}

func (t *Trie) StartWith(prefix string) bool {
	cur := t.root
	for i := 0; i < len(prefix); i++ {
		if _, ok := cur.next[prefix[i]]; !ok {
			return false
		}
		cur = cur.next[prefix[i]]
	}

	return true
}

func (t *Trie) String() string {
	// res := []string
	// for _, no := range t.root.next {
	// 	var word []rune
	// 	cur := no
	// 	for !cur.isWord {
	// 		word = append(word, cur.)
	// 	}
	// }
	return ""
}
