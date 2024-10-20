package stack

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

type Stack struct {
	data []*entity
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Pop() *entity {
	res := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]
	return res
}

func (s *Stack) Push(e *entity) {
	s.data = append(s.data, e)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

type entity struct {
	no int
	s  string
}

func Test394(t *testing.T) {
	s := "3[a]2[bc]"
	fmt.Println(decodeString(s))
}

func decodeString(s string) string {
	var (
		res   strings.Builder
		multi strings.Builder
		stk   = &Stack{data: make([]*entity, 0)}
	)

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case unicode.IsDigit(rune(c)):
			multi.WriteByte(c)
		case unicode.IsLetter(rune(c)):
			res.WriteByte(c)
		case c == '[':
			no, _ := strconv.Atoi(multi.String())
			stk.Push(&entity{
				no: no,
				s:  res.String(),
			})
			multi.Reset()
			res.Reset()
		case c == ']':
			e := stk.Pop()
			fmt.Println(e)
			tmp := strings.Builder{}
			for j := 0; j < e.no; j++ {
				tmp.WriteString(res.String())
			}
			str := e.s + tmp.String()
			res.Reset()
			res.WriteString(str)
		}
	}

	return res.String()
}

func Test394d(t *testing.T) {
	fmt.Println(decodeStringD("3[a2[c]]"))
}

func decodeStringD(s string) string {
	res, _ := dfs(s, 0)
	return res
}

func dfs(s string, i int) (string, int) {
	var (
		res   string
		multi int
	)

	for i < len(s) {
		switch {
		case unicode.IsDigit(rune(s[i])):
			multi = multi*10 + int(s[i]-'0')
		case unicode.IsLetter(rune(s[i])):
			res += string(s[i])
		case s[i] == '[':
			tmp, newIndex := dfs(s, i+1)
			for j := 0; j < multi; j++ {
				res += tmp
			}
			i = newIndex
		case s[i] == ']':
			return res, i
		}
		i++
	}

	return res, i
}
