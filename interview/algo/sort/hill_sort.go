package sort

func insertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		target := nums[i]
		j := i
		for j-1 >= 0 && nums[j-1] >= target {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = target
	}
}
