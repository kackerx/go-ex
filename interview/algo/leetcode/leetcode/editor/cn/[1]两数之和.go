package main

// leetcode submit region begin(Prohibit modification and deletion)
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

// leetcode submit region end(Prohibit modification and deletion)
