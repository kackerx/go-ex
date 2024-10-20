package main
//leetcode submit region begin(Prohibit modification and deletion)
type TrieNode struct {
	children [26]*TrieNode
	passCnt  int
	end      bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{
		children: [26]*TrieNode{},
	}}
}

func (this *Trie) getNode(s string) *TrieNode {
	cur := this.root
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if cur.children[c] == nil {
			return nil
		}
		cur = cur.children[c]
	}
	return cur
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if cur.children[c] == nil {
			cur.children[c] = &TrieNode{}
		}

		cur = cur.children[c]
	}
	cur.end = true
}

func (this *Trie) Search(word string) bool {
	node := this.getNode(word)
	return node != nil && node.end
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.getNode(prefix) != nil
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
