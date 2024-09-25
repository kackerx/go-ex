package tree

type TrieNode struct {
	childrens [26]*TrieNode
	passCnt   int
	end       bool
}

func NewTrieNode(end bool) *TrieNode {
	return &TrieNode{end: end}
}

type Triee struct {
	root *TrieNode
}

func NewTriee() *Triee {
	return &Triee{root: NewTrieNode(false)}
}

func (t *Triee) Contains(word string) bool {
	no := t.getNode(word)
	return no != nil && no.end
}

func (t *Triee) getNode(word string) *TrieNode {
	cur := t.root
	for _, ch := range word {
		if cur.childrens[ch-'a'] == nil {
			return nil
		}

		cur = cur.childrens[ch-'a']
	}

	return cur
}

func (t *Triee) HasPrefix(prefix string) bool {
	return t.getNode(prefix) != nil
}

func (t *Triee) PassCnt(prefix string) int {
	no := t.getNode(prefix)
	if no == nil {
		return 0
	}

	return no.passCnt
}

func (t *Triee) Insert(word string) {
	if t.Contains(word) {
		return
	}

	cur := t.root
	for _, ch := range word {
		if cur.childrens[ch-'a'] == nil {
			cur.childrens[ch-'a'] = NewTrieNode(false)
		}
		cur.passCnt++
		cur = cur.childrens[ch-'a']
	}

	cur.end = true
}

func (t *Triee) Remove(word string) bool {
	if !t.Contains(word) {
		return false
	}

	cur := t.root
	for _, ch := range word {
		cur.childrens[ch-'a'].passCnt--
		if cur.childrens[ch-'a'].passCnt == 0 {
			cur.childrens[ch-'a'] = nil
			return true
		}
		cur = cur.childrens[ch-'a']
	}

	cur.end = false
	return true
}
