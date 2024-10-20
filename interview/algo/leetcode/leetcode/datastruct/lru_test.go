package datastruct

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
	fmt.Println(l.Get(3))
}
