package datastruct

type TrieNode struct {
	children [26]*TrieNode
	passCnt  int
	end      bool
}

type Trie struct {
	root *TrieNode
}

func ConstructorTrie() Trie {
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

		cur.children[c].passCnt++
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

func (this *Trie) Remove(word string) bool {
	cur := this.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		cur.children[c].passCnt--
		if cur.children[c].passCnt == 0 {
			cur.children[c] = nil
			return true
		}
		cur = cur.children[c]
	}

	cur.end = false
	return true
}

func (this *Trie) PassCnt(prefix string) int {
	node := this.getNode(prefix)
	if node == nil {
		return 0
	}

	return node.passCnt
}
