package main

import "strconv"

// leetcode submit region begin(Prohibit modification and deletion)
func evalRPN(tokens []string) int {
	st := &stock{}
	opt := map[string]func(a, b int) int{
		"+": func(a, b int) int {
			return b + a
		},
		"-": func(a, b int) int {
			return b - a
		},
		"*": func(a, b int) int {
			return b * a
		},
		"/": func(a, b int) int {
			return b / a
		},
	}

	for _, token := range tokens {
		if operator, ok := opt[token]; ok {
			st.push(operator(st.pop(), st.pop()))
			continue
		}
		val, _ := strconv.Atoi(token)
		st.push(val)
	}

	return st.pop()
}

type stock []int

func (s *stock) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stock) pop() (e int) {
	e = (*s)[len(*s)-1]
	*s = (*s)[:len((*s))-1]
	return
}

func (s *stock) isEmpty() bool {
	return len(*s) == 0
}

func (s *stock) push(e int) {
	*s = append(*s, e)
}

// leetcode submit region end(Prohibit modification and deletion)
