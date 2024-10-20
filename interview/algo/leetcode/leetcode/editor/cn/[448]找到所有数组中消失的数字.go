package main

// leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)

	for _, num := range nums {
		if nums[abs(num)-1] > 0 {
			nums[abs(num)-1] *= -1
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// leetcode submit region end(Prohibit modification and deletion)
