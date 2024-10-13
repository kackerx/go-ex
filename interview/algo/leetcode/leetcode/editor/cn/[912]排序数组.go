package main

// leetcode submit region begin(Prohibit modification and deletion)
func sortArray(nums []int) []int {
	tmp := make([]int, len(nums))
	mergeSort1(nums, tmp, 0, len(nums)-1)
	return nums
}

func mergeSort1(nums, tmp []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	mergeSort1(nums, tmp, lo, mid)
	mergeSort1(nums, tmp, mid+1, hi)
	merge2(nums, tmp, lo, mid, hi)
}

func merge2(nums, tmp []int, lo, mid, hi int) {
	copy(tmp[lo:hi+1], nums[lo:hi+1])

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			nums[k] = tmp[j]
			j++
		} else if j > hi || tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
