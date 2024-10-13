package main

// leetcode submit region begin(Prohibit modification and deletion)
type pair struct {
	index, val int
}

func countSmaller(nums []int) []int {
	n := len(nums)
	count := make([]int, n)
	arr := make([]*pair, n)
	tmp := make([]*pair, n)

	for i, v := range nums {
		arr[i] = &pair{index: i, val: v}
	}

	mergeSort(arr, tmp, count, 0, n-1)
	return count
}

func mergeSort(arr, tmp []*pair, count []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	mergeSort(arr, tmp, count, lo, mid)
	mergeSort(arr, tmp, count, mid+1, hi)

	merge(arr, tmp, count, lo, mid, hi)
}

func merge(arr, tmp []*pair, count []int, lo, mid, hi int) {
	copy(tmp[lo:hi+1], arr[lo:hi+1])

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = tmp[j]
			j++
		} else if j > hi {
			arr[k] = tmp[i]
			count[arr[k].index] += j-(mid+1)
			i++
		} else if tmp[i].val <= tmp[j].val {
			arr[k] = tmp[i]
			count[arr[k].index] += j-(mid+1)
			i++
		} else {
			arr[k] = tmp[j]
			j++
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
