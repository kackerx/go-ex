package leetcode

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

var (
	arr = []int{1, 8, 5, 29, 1, 10, 5, 43, 38, 29, 4, 3, 8, 9, 10, 17}
)

func TestIsValid(t *testing.T) {
	s := "[](()){{}}"
	fmt.Println(IsValid(s))
}

func Test75(t *testing.T) {
	a := []int{1, 1, 2, 2, 0, 1, 2, 0, 1, 2, 1}
	SortColors(a)

	fmt.Println(a)
}

func Test567(t *testing.T) {
	fmt.Println(checkInclusion("eidbaooo", "ab"))
}

func checkInclusion(s1 string, s2 string) bool {
	var left, right, valid int

	win := make(map[rune]int)
	need := make(map[rune]int)

	for _, i := range s2 {
		need[i]++
	}

	for right < len(s1) {
		c := s1[right]
		right++

		if v, ok := need[rune(c)]; ok {
			win[rune(c)]++
			if win[rune(c)] == v {
				valid++
			}
		}

		if right-left >= len(s2) {
			if valid == len(need) {
				return true
			}

			d := s1[left]
			left++
			if v, ok := need[rune(d)]; ok {
				if win[rune(d)] == v {
					valid--
				}
				win[rune(d)]--
			}
		}
	}

	return false
}

func Test3(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	var left, right, res int
	win := make(map[byte]int)

	for right < len(s) {
		c := s[right]
		right++
		win[c]++

		if win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}

		res = max(res, right-left)

	}

	return res
}

func Test438(t *testing.T) {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
	need, win := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	var left, right, valid int

	var res []int
	for right < len(s) {
		c := s[right]
		right++

		if v, ok := need[c]; ok {
			win[c]++
			if win[c] == v {
				valid++
			}

		}

		if right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}

			d := s[left]
			left++
			if v, ok := need[d]; ok {
				if win[d] == v {
					valid--
				}
				win[d]--
			}
		}
	}

	return res
}

func TestCeil(t *testing.T) {
	fmt.Println(leftBound([]int{1, 2, 2, 2, 3, 3}, 3))
}

func leftBound(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if l >= len(nums) {
		return -1
	}
	return l
}

func TestJz20(t *testing.T) {
	GetLeastNumbers(arr, 5)
}

func Test225(t *testing.T) {
	s := NewMyStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

func TestName(t *testing.T) {
	q := NewMyQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}

func Test167(t *testing.T) {
}

func TwoSum(numbers []int, target int) []int {
	// left := 0
	// right := len(numbers) - 1
	// for left < right {
	// 	sum := numbers[left] + numbers[right]
	// 	if sum == target {
	// 		return []int{left + 1, right + 1}
	// 	} else if sum < target {
	// 		left++
	// 	} else {
	// 		right--
	// 	}
	// }
	//
	// return []int{-1, -1}

	m := make(map[int]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		m[numbers[i]] = i
	}

	fmt.Println(m)

	for k, v := range m {
		otherK := target - k
		if otherV, ok := m[otherK]; ok {
			return []int{v + 1, otherV + 1}
		}
	}

	return []int{-1, -1}
}

func TestScores(t *testing.T) {
	scores := []int{1, 3, 5, 2, 4, 5, 2, 3, 3, 4, 5, 0, 0}

	count := make([]int, 5+1)
	for _, score := range scores {
		count[score]++
	}

	fmt.Println(count)

	for i := 1; i < len(count); i++ {
		count[i] = count[i] + count[i-1]
	}

	fmt.Println(count)
}

func Test51(t *testing.T) {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	var res [][]string
	board := newBoard(n)

	backTrack(board, 0, &res)

	return res
}

func backTrack(board [][]byte, row int, res *[][]string) {
	if row == len(board) {
		var item []string
		for _, bytes := range board {
			item = append(item, string(bytes))
		}
		*res = append(*res, item)
	}

	n := len(board)
	for col := 0; col < n; col++ {
		if !isValid(board, row, col) {
			continue
		}

		board[row][col] = 'Q'
		backTrack(board, row+1, res)
		board[row][col] = '.'
	}
}

func isValid(board [][]byte, row int, col int) bool {
	n := len(board)
	// 检测列上是否有Q
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 检测右上方斜列是否有
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// 检测左上方斜列是否有
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func newBoard(n int) [][]byte {
	res := make([][]byte, n)
	for i := 0; i < n; i++ {
		res[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			res[i][j] = '.'
		}
	}
	return res
}

func Test78(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	var res [][]int
	backtrackk(nums, 0, []int{}, &res)

	return res
}

func backtrackk(nums []int, index int, track []int, res *[][]int) {
	*res = append(*res, slices.Clone(track))

	for i := index; i < len(nums); i++ {
		// if slices.Contains(track, nums[i]) {
		// 	continue
		// }

		track = append(track, nums[i])
		backtrackk(nums, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func Test77(t *testing.T) {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	var res [][]int
	back(n, k, 1, []int{}, &res)
	return res
}

func back(n, k, index int, track []int, res *[][]int) {
	// if len(track) == k {
	*res = append(*res, slices.Clone(track))
	// return
	// }

	for i := index; i < n+1; i++ {
		if slices.Contains(track, i) {
			continue
		}

		track = append(track, i)
		back(n, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func Test543(t *testing.T) {
	fmt.Println(diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}))
}

func Test46(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	backtrack(nums, track, res)

	return res
}

func backtrack(nums []int, track []int, res [][]int) {
	if len(track) == len(nums) {
		res = append(res, track)
		return
	}

	for i := 0; i < len(nums); i++ {
		if slices.Contains(track, nums[i]) {
			continue
		}

		track = append(track, nums[i])
		backtrack(nums, track, res)
		track = track[:len(track)-1]
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDiameter = 0

func diameterOfBinaryTree(root *TreeNode) int {
	// 记录最大直径的长度

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := maxDepth(root.Left)
		rightMax := maxDepth(root.Right)
		// 后序位置，顺便计算最大直径
		myDiameter := leftMax + rightMax
		maxDiameter = max(maxDiameter, myDiameter)
		return 1 + max(leftMax, rightMax)
	}

	maxDepth(root)

	return maxDiameter
}

// func Test739(t *testing.T) {
// 	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 69, 72, 76, 73}))
// }
//
// type stack []int
//
// func (s *stack) peek() int {
// 	return (*s)[len(*s)-1]
// }
//
// func (s *stack) pop() (e int) {
// 	e = (*s)[len(*s)-1]
// 	*s = (*s)[:len((*s))-1]
// 	return
// }
//
// func (s *stack) isEmpty() bool {
// 	return len(*s) == 0
// }
//
// func (s *stack) push(e int) {
// 	*s = append(*s, e)
// }
//
// func dailyTemperatures(temperatures []int) []int {
// 	res := make([]int, len(temperatures))
//
// 	st := &stack{}
//
// 	for i := 0; i < len(temperatures); i++ {
// 		for !st.isEmpty() {
// 			if temperatures[i] > temperatures[st.peek()] { // 一直比较当前遍历值和栈顶如果比栈顶大, 就说明当前值是相比栈顶的下标的值第一个大的, 出栈记录结果
// 				index := st.pop()
// 				res[index] = i - index
// 			} else { // 循环直到碰到当前值小于等于栈顶或者栈为空
// 				break
// 			}
// 		}
//
// 		st.push(i) // 当前值入栈
// 	}
//
// 	return res
// }

func TestKnapsack(t *testing.T) {
	fmt.Println(knapsack(4, 3, []int{2, 1, 3}, []int{4, 2, 3}))
}

func knapsack(W, N int, wt, val []int) int {
	dp := make([][]int, N+1) // 前i个物品装进去且重量为w时的最大价值
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-wt[i-1] < 0 { // 如果当前重量不足以减去物品i的重量, 直接是什么都不装的i-1时的价值
				dp[i][w] = dp[i-1][w]
				continue
			}
			dp[i][w] = max(
				dp[i-1][w-wt[i-1]]+val[i-1], // 把i装进去, 当前重量-i的重量加上i的价值
				dp[i-1][w],                  // 不把i装进去, i-1的价值, 重量不变
			)
		}
	}

	return dp[N][W]
}

func Test69(t *testing.T) {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	l := 0
	r := x + 1
	for l < r {
		mid := l + (r-l+1)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l
}

func Test111(t *testing.T) {
	root := &TreeNode{3, &TreeNode{Val: 9}, &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

	fmt.Println(minDept(root))
}

func minDept(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDept(root.Left)
	right := minDept(root.Right)

	return min(left, right) + 1
}

type queu []*TreeNode

func (q *queu) Get() *TreeNode {
	res := (*q)[0]
	*q = append((*q)[:0], (*q)[1:]...)
	return res
}

func (q *queu) Put(no *TreeNode) {
	*q = append(*q, no)
}

func (q *queu) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queu) Len() int {
	return len(*q)
}

func minDepth(root *TreeNode) int {
	var (
		q     = &queu{}
		depth = 1
	)

	q.Put(root)
	for !q.IsEmpty() {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			no := q.Get()
			if no.Left != nil {
				q.Put(no.Left)
			}
			if no.Right != nil {
				q.Put(no.Right)
			}
			if no.Left == nil && no.Right == nil {
				return depth
			}
		}
		depth++
	}

	return depth
}

func Test151(t *testing.T) {
	s := " he   world  "

	bs := cleanSpace(s)
	rev(bs, 0, len(bs)-1)

	for i := 0; i < len(bs); {
		for j := i; j < len(bs); j++ {
			if j+1 == len(bs) || bs[j+1] == ' ' {
				rev(bs, i, j)
				i = j + 2
				break
			}
		}
	}

	fmt.Println(string(bs))
}

func cleanSpace(s string) []byte {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			bs = append(bs, s[i])
		} else if len(bs) > 0 && bs[len(bs)-1] != ' ' {
			bs = append(bs, ' ')
		}
	}

	if bs[len(bs)-1] == ' ' {
		bs = bs[:len(bs)-1]
	}

	return bs
}

func rev(bs []byte, i, j int) {
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
}

// ---

func reverseWords(s string) string {
	var sb strings.Builder
	// 先清洗一下数据，把多于的空格都删掉
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ' ' {
			// 单词中的字母/数字
			sb.WriteByte(c)
		} else if sb.Len() > 0 && sb.String()[sb.Len()-1] != ' ' {
			// 单词之间保留一个空格
			sb.WriteByte(' ')
		}
	}
	if sb.Len() == 0 {
		return ""
	}
	// 末尾如果有空格，清除之
	cleaned := sb.String()
	if cleaned[len(cleaned)-1] == ' ' {
		cleaned = cleaned[:len(cleaned)-1]
	}

	// 清洗之后的字符串
	chars := []byte(cleaned)
	n := len(chars)
	// 进行单词的翻转，先整体翻转
	reverse(chars, 0, n-1)
	// 再把每个单词翻转
	for i := 0; i < n; {
		for j := i; j < n; j++ {
			if j+1 == n || chars[j+1] == ' ' {
				// chars[i..j] 是一个单词，翻转之
				reverse(chars, i, j)
				// 把 i 置为下一个单词的首字母
				i = j + 2
				break
			}
		}
	}
	// 最后得到题目想要的结果
	return string(chars)
}

// 翻转 arr[i..j]
func reverse(arr []byte, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func Test209(t *testing.T) {
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

func minSubArrayLen(target int, nums []int) int {
	var left, right, sum int
	minLen := math.MaxInt

	for right < len(nums) {
		val := nums[right]
		right++
		sum += val

		for sum >= target {
			minLen = min(minLen, right-left)
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt {
		return 0
	} else {
		return minLen
	}
}

type User struct {
	Name string
}

func (u User) name() {
	return
}

func TestAr(t *testing.T) {
	u := &User{}
	u.name()
}
