package leetcode

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"testing"
)

var head = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

var head2 = &ListNode{Val: 2, Next: &ListNode{Val: 7, Next: &ListNode{Val: 9}}}

var head3 = &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 8, Next: &ListNode{Val: 7, Next: &ListNode{Val: 9}}}}}

func Test148(t *testing.T) {
	// sortList(head)
	// print(merge(head, head2))
	print(sortList(head3))
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := GetMid(head)
	left := head
	right := mid.Next
	mid.Next = nil

	leftList := sortList(left)
	rightList := sortList(right)

	return merge(leftList, rightList)
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	p, p1, p2 := dummy, left, right

	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy

	p1 := list1
	p2 := list2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}

func GetMid(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func TestFoo(t *testing.T) {
	s := "aaabbccccdd"
	res := make([]int, len(s))

	for i := range s {
		res[s[i]-97]++
	}

	fmt.Println(res)
}

func TestBar(t *testing.T) {
	fmt.Println(tiao(3))
}

func tiao(n int) int {
	if n <= 2 {
		return n
	}

	return tiao(n-1) + tiao(n-2)
}

func TestStack(t *testing.T) {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	reve2(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

func getLast(s *Stack[int]) int {
	if s.Len() == 1 {
		return s.Pop()
	}

	temp := s.Pop()
	res := getLast(s) // 逆天
	s.Push(temp)

	return res
}

func reve2(s *Stack[int]) *Stack[int] {
	if s.Len() < 2 {
		return s
	}

	last := getLast(s)
	s = reve2(s)
	s.Push(last)

	return s
}

func reve(s *Stack[int]) *Stack[int] {
	if s.Len() < 2 {
		return s
	}

	top := s.Pop()
	s = reve(s)
	return pushFirst(s, top)
}

func pushFirst(s *Stack[int], e int) *Stack[int] {
	if s.IsEmpty() {
		s.Push(e)
		return s
	}

	temp := s.Pop() // 前序弹出
	s = pushFirst(s, e)
	s.Push(temp) // 后序弹入

	return s
}

func print(head *ListNode) {
	if head == nil {
		return
	}

	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func TestReveN(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	// print(ReveN(head, 3))
	// print(reverseBetween(head, 2, 4))
	print(reverseBetweenR(head, 2, 4))
	// print(ReveNR(head, 3))
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return ReveN(head, right)
	}

	pre := head
	for i := 0; i < left-1-1; i++ {
		pre = pre.Next
	}

	pre.Next = ReveN(pre.Next, right-left+1)
	return head
}

func reverseBetweenR(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return ReveN(head, right)
	}

	head.Next = reverseBetweenR(head.Next, left-1, right-1)
	return head
}

func ReveNR(head *ListNode, n int) *ListNode {
	if head == nil || n == 1 {
		return head
	}

	next := head
	for i := 0; i < n; i++ {
		next = next.Next
	}

	newHead := ReveNR(head.Next, n-1)
	head.Next.Next = head
	head.Next = next
	return newHead
}

func ReveN(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre, cur, next *ListNode
	cur = head
	for i := 0; i < n; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	head.Next = cur
	return pre
}

func Test25(t *testing.T) {

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}

		b = b.Next
	}

	newHead := ReveN(a, k)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func Test61(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	print(rotateRight(head, 2))

}

func rotateRight(head *ListNode, k int) *ListNode {
	n := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}

	cur.Next = head
	for i := 0; i < (n-k)%n; i++ {
		cur = cur.Next
	}

	newHead := cur.Next
	cur.Next = nil

	return newHead
}

type pair struct {
	num, cnt int
}

type pq []pair

func (h *pq) Len() int {
	return len(*h)
}

func (h *pq) Less(i, j int) bool {
	return (*h)[i].cnt < (*h)[j].cnt
}

func (h *pq) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *pq) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *pq) Pop() any {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h *pq) Peek() int {
	return (*h)[0].cnt
}

func Test347(t *testing.T) {
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))
	// 4: 1
	// 1: 1
	// -1: 2
	// 2: 2
	// 3: 1
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	q := &pq{}
	heap.Init(q)

	for key, val := range m {
		if q.Len() < k {
			heap.Push(q, pair{
				num: key,
				cnt: val,
			})
			continue
		}

		if m[q.Peek()] <= val {
			heap.Pop(q)
			heap.Push(q, pair{
				num: key,
				cnt: val,
			})
		}
	}

	res := make([]int, 0)

	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(q).(pair).num)
	}

	return res
}

func Test1475(t *testing.T) {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}

func finalPrices(prices []int) []int {
	st := Stack[int]{}
	res := make([]int, len(prices))

	for i := 0; i < len(prices); i++ {
		for !st.IsEmpty() {
			if prices[i] < prices[st.Peek()] {
				topIndex := st.Pop()
				res[topIndex] = prices[topIndex] - prices[i]
			} else {
				break
			}
		}
		st.Push(i)
	}

	for !st.IsEmpty() {
		topIndex := st.Pop()
		res[topIndex] = prices[topIndex]
	}

	return res
}

func Test239(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
	// m := &MonotonicQueue{list.New()}
	// m.Push(1)
	// fmt.Println(m.Max())
	// m.Push(3)
	// fmt.Println(m.Max())
	// m.Push(-1)
	// fmt.Println(m.Max())
	// m.Push(-3)
	// fmt.Println(m.Max())
	// m.Push(5)
	// fmt.Println(m.Max())
	// m.Pop()
	// fmt.Println(m.Max())
	// m.Pop()
	// fmt.Println(m.Max())
}

type MonotonicQueue struct {
	maxq *list.List
}

func (m *MonotonicQueue) Push(e int) {
	for !m.IsEmpty() && e > m.maxq.Back().Value.(int) {
		m.maxq.Remove(m.maxq.Back())
	}
	m.maxq.PushBack(e)
}

func (m *MonotonicQueue) Pop(e int) {
	if !m.IsEmpty() && m.maxq.Front().Value.(int) == e {
		m.maxq.Remove(m.maxq.Front())
	}
}

func (m *MonotonicQueue) Max() int {
	if !m.IsEmpty() {
		return m.maxq.Front().Value.(int)
	}

	return -1
}

func (m *MonotonicQueue) IsEmpty() bool {
	return m.maxq.Len() == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	win := &MonotonicQueue{list.New()}
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			win.Push(nums[i])
			continue
		}

		win.Push(nums[i])
		res = append(res, win.Max())
		win.Pop(nums[i-k+1])
	}

	return res
}

func maxSlidingWindowx(nums []int, k int) []int {
	left, right := 0, k
	res := make([]int, 0)

	for right <= len(nums) {
		res = append(res, slices.Max(nums[left:right]))
		left++
		right++
	}

	return res
}

func Test622(t *testing.T) {
	// [1,3,2,5,3,null,9]
	fmt.Println(widthOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}))
}

type Pair struct {
	node *TreeNode
	id   int
}

type queue1 []*Pair

func (q *queue1) Get() *Pair {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *queue1) Put(no *Pair) {
	*q = append(*q, no)
}

func (q *queue1) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue1) Len() int {
	return len(*q)
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return -1
	}

	var maxWidth, left, right int
	q := &queue1{}
	q.Put(&Pair{root, 1})

	for !q.IsEmpty() {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			no := q.Get()
			if i == 0 {
				left = no.id
			}
			if i == sz-1 {
				right = no.id
			}

			if no.node.Left != nil {
				q.Put(&Pair{no.node.Left, 2 * (no.id)})
			}

			if no.node.Right != nil {
				q.Put(&Pair{no.node.Right, 2*(no.id) + 1})
			}
		}
		maxWidth = max(maxWidth, right-left+1)
	}

	return maxWidth
}

// ---
func Test297(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   02,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	codec := Constructor()
	serialize := codec.serialize(root)

	res := codec.deserialize(serialize)
	fmt.Println(res)
}

type Codec struct {
	sep  string
	null string
	res  []string
}

func Constructor() Codec {
	return Codec{",", "#", make([]string, 0)}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// sb := strings.Builder{}
	this.traverse(root)
	res := strings.Join(this.res, this.sep)
	return res
}

func (this *Codec) traverse(no *TreeNode) {
	if no == nil {
		this.res = append(this.res, this.null)
		return
	}

	this.res = append(this.res, strconv.Itoa(no.Val))
	this.traverse(no.Left)
	this.traverse(no.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	return this.Deserializes(&this.res)
}

func (this *Codec) Deserializes(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}

	first := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if first == this.null {
		return nil
	}

	val, _ := strconv.Atoi(first)
	no := &TreeNode{Val: val}

	no.Left = this.Deserializes(nodes)
	no.Right = this.Deserializes(nodes)

	return no
}

func Test113(t *testing.T) {
	fmt.Println(pathSum(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}, 19))
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	traverse(root, targetSum, []int{}, &res)

	return res
}

func traverse(no *TreeNode, target int, track []int, res *[][]int) {
	if no == nil {
		return
	}

	track = append(track, no.Val)

	if no.Left == no.Right && target == no.Val {
		*res = append(*res, slices.Clone(track))
		return
	}

	traverse(no.Left, target-no.Val, track, res)
	traverse(no.Right, target-no.Val, track, res)
	track = track[:len(track)-1]
}

func Test256(t *testing.T) {
	fmt.Println(lowestCommonAncestor(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}, &TreeNode{Val: 5}, &TreeNode{Val: 1}))
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return find(root, p.Val, q.Val)
}

func find(no *TreeNode, val1, val2 int) *TreeNode {
	if no == nil {
		return nil
	}

	// 前序 (先遇到的一定是公共祖先
	if no.Val == val1 || no.Val == val2 {
		return no
	}

	left := find(no.Left, val1, val2)
	right := find(no.Right, val1, val2)

	// 后续位置 (左子树和右子树都不为nil那就是当前节点
	if left != nil && right == nil {
		return no
	}

	if left != nil {
		return left
	} else {
		return right
	}
}

func Test376(t *testing.T) {
	d := NewDiff([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8})
	diff := d.diff

	// Remove zero differences
	filteredDiff := []int{}
	for _, v := range diff {
		if v != 0 {
			filteredDiff = append(filteredDiff, v)
		}
	}

	if len(filteredDiff) == 0 {
		return
	}

	count := 1 // At least one element is a valid wiggle sequence

	// Calculate the wiggle sequence length
	for i := 1; i < len(filteredDiff); i++ {
		if (filteredDiff[i-1] > 0 && filteredDiff[i] < 0) || (filteredDiff[i-1] < 0 && filteredDiff[i] > 0) {
			count++
		}
	}

	fmt.Println(count + 1) // Add 1 to include the first element in the sequence
}

type Diff struct {
	diff []int
}

func NewDiff(data []int) *Diff {
	diff := make([]int, len(data)-1)

	// diff[0] = data[0]
	for i := 0; i < len(data)-1; i++ {
		diff[i] = data[i+1] - data[i]
	}

	return &Diff{diff: diff}
}

func Test55(t *testing.T) {
	canJump([]int{2, 3, 1, 1, 4})
}

func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if farthest <= i {
			return false
		}
	}

	return farthest >= len(nums)-1
}

func Test1(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
func twoSum(nums []int, target int) []int {
	var res []int
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if val, ok := m[target-nums[i]]; !ok {
			m[nums[i]] = i
		} else {
			res = append(res, val, i)
			break
		}
	}

	return res
}

func Test77a(t *testing.T) {
	fmt.Println(combine1(4, 3))
}

func combine1(n int, k int) [][]int {
	var res [][]int
	back1(n, k, 1, []int{}, &res)
	return res
}

func back1(n, k, idx int, track []int, res *[][]int) {
	if len(track) == k {
		*res = append(*res, slices.Clone(track))
		return
	}

	for i := idx; i < n+1; i++ {
		track = append(track, i)
		back1(n, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}

func Test216(t *testing.T) {
	fmt.Println(combinationSum3(9, 45))
}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	back3(n, k, 1, 0, []int{}, &res)
	return res
}

func back3(n, k, idx, sum int, track []int, res *[][]int) {
	if len(track) == k && sum == n {
		*res = append(*res, slices.Clone(track))
		return
	}

	for i := idx; i <= 9; i++ {
		track = append(track, i)
		sum += i
		back3(n, k, i+1, sum, track, res)
		sum -= i
		track = track[:len(track)-1]
	}
}

func Test46a(t *testing.T) {
	fmt.Println(permute1([]int{1, 2, 3}))
}

func permute1(nums []int) [][]int {
	res := make([][]int, 0)
	used := map[int]bool{}
	back2(nums, []int{}, &res, used)
	return res
}

func back2(nums []int, track []int, res *[][]int, used map[int]bool) {
	if len(track) == len(nums) {
		*res = append(*res, slices.Clone(track))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		back2(nums, track, res, used)
		used[i] = false
		track = track[:len(track)-1]
	}
}

func Test15(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	var res [][]int

	slices.Sort(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		twoRes := twoSuma(nums, i+1, -nums[i])
		for _, tuple := range twoRes {
			tuple = append(tuple, nums[i])
			res = append(res, tuple)
		}

		for i+1 < n && nums[i] == nums[i+1] {
			i++
		}

	}

	return res
}

func twoSuma(nums []int, start, target int) [][]int {
	res := make([][]int, 0)
	l, r := start, len(nums)-1
	for l < r {
		left, right := nums[l], nums[r]
		sum := nums[l] + nums[r]
		if sum == target {
			res = append(res, []int{nums[l], nums[r]})
			for l < r && nums[l] == left {
				l++
			}
			for l < r && nums[r] == right {
				r--
			}
		} else if sum < target {
			for l < r && nums[l] == left {
				l++
			}
		} else {
			for l < r && nums[r] == right {
				r--
			}
		}
	}

	return res
}

func Test80(t *testing.T) {
	r := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(r))

	fmt.Println(r)
}

func removeDuplicates(nums []int) int {
	var slow, fast, count int

	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && count < 2 {
			slow++
			nums[slow] = nums[fast]
		}

		count++
		fast++
		if fast < len(nums) && nums[fast] != nums[fast-1] {
			count = 0
		}
	}

	return slow + 1
}

func Test88(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge1(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	k := len(nums1) - 1
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums1[j]
		j--
		k--
	}
}

func Test76(t *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	win, need := make(map[byte]int, 0), make(map[byte]int, 0)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	l, r, valid, start := 0, 0, 0, 0
	length := math.MaxInt
	for r < len(s) {
		c := s[r]
		r++

		if v, ok := need[c]; ok {
			win[c]++
			if win[c] == v {
				valid++
			}
		}

		for valid == len(need) {
			if r-l < length {
				start = l
				length = r - l
			}

			d := s[l]
			l++
			if v, ok := need[d]; ok {
				if win[d] == v {
					valid--
				}
				win[d]--
			}
		}
	}

	if length == math.MaxInt {
		return ""
	} else {
		return s[start : start+length]
	}
}

func Test567a(t *testing.T) {
	fmt.Println(checkInclusiona("ab", "eidbaooo"))
}

func checkInclusiona(s1 string, s2 string) bool {
	win, need := make(map[byte]int, 0), make(map[byte]int, 0)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	var l, r, valid int

	for r < len(s2) {
		c := s2[r]
		r++

		if v, ok := need[c]; ok {
			win[c]++
			if win[c] == v {
				valid++
			}
		}

		if r-l >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[l]
			l++

			if v, ok := need[d]; ok {
				if win[d] == v {
					valid--
				}
				win[d]--
			}
		}
	}

	return false
}

func Test438a(t *testing.T) {
	fmt.Println(findAnagramsa("cbaebabacd", "abc"))
}

func findAnagramsa(s string, p string) []int {
	win, need := make(map[byte]int, 0), make(map[byte]int, 0)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	var res []int
	var l, r, valid int

	for r < len(s) {
		c := s[r]
		r++

		if v, ok := need[c]; ok {
			win[c]++
			if win[c] == v {
				valid++
			}
		}

		for r-l >= len(p) {
			if valid == len(need) {
				res = append(res, l)
			}
			d := s[l]
			l++

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

func Test3a(t *testing.T) {
	fmt.Println(lengthOfLongestSubstringa("pwwkew"))
}

func lengthOfLongestSubstringa(s string) int {
	var l, r, res int
	win := make(map[byte]int, 0)

	for r < len(s) {
		c := s[r]
		r++

		win[c]++

		for win[c] > 1 {
			d := s[l]
			win[d]--
			l++
		}

		res = max(res, r-l)

	}

	return res
}

func Test255(t *testing.T) {
	s := Constructo()
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Top())
}

type MyStac struct {
	s1, s2 []int
}

func Constructo() MyStac {
	return MyStac{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

func (this *MyStac) Push(x int) {
	this.s1 = append(this.s1, x)
	for len(this.s2) != 0 {
		val := this.s2[0]
		this.s2 = this.s2[1:len(this.s2)]
		this.s1 = append(this.s1, val)
	}

	this.s2, this.s1 = this.s1, this.s2
}

func (this *MyStac) Pop() int {
	val := this.s2[0]
	this.s2 = this.s2[1:len(this.s2)]
	return val
}

func (this *MyStac) Top() int {
	return this.s2[0]
}

func (this *MyStac) Empty() bool {
	return len(this.s2) == 0
}

func Test71(t *testing.T) {
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))
}

func simplifyPath(path string) string {
	s := &Stack[string]{}
	pathList := strings.Split(path, "/")
	slices.Reverse(pathList)
	for _, item := range pathList {
		if item == "." || item == "" {
			continue
		}

		s.Push(item)
	}

	s2 := &Stack[string]{}
	for !s.IsEmpty() {
		v := s.Pop()
		if v == ".." {
			if s2.IsEmpty() {
				continue
			}
			s2.Pop()
			continue
		}

		s2.Push(v)
	}

	var res string
	for !s2.IsEmpty() {
		res = fmt.Sprintf("/%s", s2.Pop()) + res
	}

	return res
}

func TestCalculate(t *testing.T) {
	fmt.Println(calculateGreaterElement([]int{2, 1, 2, 4, 3}))
}

func calculateGreaterElement(nums []int) []int {
	res := make([]int, len(nums))
	s := &Stack[int]{}

	for i := len(nums) - 1; i >= 0; i-- {
		for !s.IsEmpty() && s.Peek() <= nums[i] {
			s.Pop()
		}

		if s.IsEmpty() {
			res[i] = -1
		} else {
			res[i] = s.Peek()
		}

		s.Push(nums[i])
	}

	return res
}

func Test128(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
	s := &set{}
	for _, num := range nums {
		s.add(num)
	}

	var maxLen int

	for num := range *s {
		if s.has(num - 1) {
			continue
		}

		curNum := num
		curCnt := 1
		for {
			if s.has(curNum + 1) {
				curNum++
				curCnt++
			} else {
				break
			}
		}
		maxLen = max(maxLen, curCnt)
	}

	return maxLen
}

type set map[int]struct{}

func (s *set) has(key int) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *set) add(key int) {
	(*s)[key] = struct{}{}
}
