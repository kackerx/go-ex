package string

import (
	"math"
	"testing"
)

// ****** 在最低位添加一个数字 ******
// number 的进制
// int R = 4;
// 想在 number 的最低位添加的数字
// int appendVal = 0~3 中的任意数字;
// 运算，在最低位添加一位
// number = R * number + appendVal;

// ****** 在最高位删除一个数字 ******
// number 的进制
// int R = 4;
// number 最高位的数字
// int removeVal = 0~3 中的任意数字;
// 此时 number 的位数
// int L = ?;
// 运算，删除最高位数字
// number = number - removeVal * R^(L-1);

func Test187(t *testing.T) {

}

func findRepeatedDnaSequences(s string) []string {
	L := 10 // 数字位个数
	R := 10 // 进制
	RL := int(math.Pow(float64(R), float64(L-1)))

	// 编码字符串为一个10位数字
	nums := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			nums[i] = 0
		case 'C':
			nums[i] = 1
		case 'G':
			nums[i] = 2
		case 'T':
			nums[i] = 3
		}
	}

	seen := NewSet()
	resMap := make(map[string]struct{})
	left, right, win := 0, 0, 0
	for right < len(nums) {
		// 增大窗口: 最低位添加数字
		v := nums[right]
		win = win*R + v
		right++

		if right-left == L {
			if seen.Has(win) {
				resMap[s[left:right]] = struct{}{}
			} else {
				seen.Add(win)
			}

			// 缩小窗口: 最高位缩小窗口
			v = nums[left]
			win -= v * RL
		}
	}

	res := make([]string, 0, len(resMap))
	for k := range resMap {
		res = append(res, k)
	}

	return res
}

type Set struct {
	m map[int]struct{}
}

func NewSet() *Set {
	return &Set{make(map[int]struct{})}
}

// func (s *Set) Keys() []string {
// 	res := make([]string, 0, len(s.m))
// 	for k, _ := range s.m {
// 		res = append(res, k)
// 	}
//
// 	return res
// }

func (s *Set) Add(e int) {
	s.m[e] = struct{}{}
}

func (s *Set) Has(e int) bool {
	_, exist := s.m[e]
	return exist
}
