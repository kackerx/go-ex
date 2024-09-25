package leetcode

import (
	"go-bobo/interview/algo/struct/stack"
)

func IsValid(s string) bool {
	stock := stack.NewArrayStock[string](10)
	for _, item := range s {
		switch string(item) {
		case "[", "(", "{":
			stock.Push(string(item))
		case "]":
			if stock.Pop() != "[" {
				return false
			}
		case ")":
			if stock.Pop() != "(" {
				return false
			}
		case "}":
			if stock.Pop() != "{" {
				return false
			}

		}
	}

	return stock.IsEmpty()
}
