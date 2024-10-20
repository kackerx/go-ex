package datastruct

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := ConstructorTrie()
	trie.Insert("apple")
	trie.Insert("appl")
	trie.Insert("applee")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("applee"))
	fmt.Println(trie.PassCnt("apple"))

	trie.Remove("apple")
	fmt.Println(trie.PassCnt("apple"))

	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("applee"))
	// fmt.Println(trie.Search("apple"))
	// fmt.Println(trie.StartsWith("a"))
}

func TestFoo(t *testing.T) {
	a := TrieNode{}
	fmt.Println(len(a.children))
}
