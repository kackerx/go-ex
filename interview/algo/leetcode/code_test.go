package leetcode

import (
	"fmt"
	"testing"
)

var head = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

var head2 = &ListNode{Val: 2, Next: &ListNode{Val: 7, Next: &ListNode{Val: 9}}}

var head3 = &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 8, Next: &ListNode{Val: 7, Next: &ListNode{Val: 9}}}}}

func Test151a(t *testing.T) {
	s := "the sky is blue"
	fmt.Println(reverseWordsa(s))
}

func reverseWordsa(s string) string {
	bs := []byte(s)

	bs = clearSpace(bs)
	reve(bs, 0, len(bs)-1)

	for i := 0; i < len(bs); {
		for j := i; j < len(bs); j++ {
			if j == len(bs)-1 {
				reve(bs, i, j)
				i = j + 1
			} else if bs[j] == ' ' {
				reve(bs, i, j-1)
				i = j + 1
			}
		}
	}

	return string(bs)
}

func reve(bs []byte, start, end int) {
	for start < end {
		bs[start], bs[end] = bs[end], bs[start]
		start++
		end--
	}
}

func clearSpace(bs []byte) []byte {
	res := make([]byte, 0)
	for i := 0; i < len(bs); i++ {
		if i == 0 && bs[i] == ' ' {
			continue
		}

		if bs[i] == ' ' && bs[i-1] == ' ' {
			continue
		}

		res = append(res, bs[i])
	}

	if res[len(res)-1] == ' ' {
		res = res[:len(res)-1]
	}

	return res
}

func TestAA(t *testing.T) {
	fmt.Printf("%d\n", 99-'a')
}
