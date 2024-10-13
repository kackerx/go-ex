package sort

import (
	"fmt"
	"testing"
)

var nums = []int{1, 7, 5, 3, 8, 9, 10, 13, 2, 7}

func TestMaopao(t *testing.T) {
	// maopao1(nums)
	// select1(nums)
	// quick1(nums, 0, len(nums)-1)
	// insert1(nums)
	// mergeSort(nums, 0, len(nums)-1)
	// sort(nums)
	sort1(nums)

	fmt.Println(nums)
}

func maopao1(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		var flag bool
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j+1] < nums[j] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func select1(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
}

func insert1(nums []int) {
	for i := 0; i < len(nums); i++ {
		target := nums[i]
		j := i
		for j-1 >= 0 && nums[j-1] > target {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = target
	}
}

func quick1(nums []int, left, right int) {
	if left >= right {
		return
	}

	i := left + 1
	lt := left // [left+1, i-1]
	gt := right + 1
	target := nums[left]

	for i < gt {
		if nums[i] == target {
			i++
		} else if nums[i] < target {
			lt++
			nums[lt], nums[i] = nums[i], nums[lt]
			i++
		} else {
			gt--
			nums[gt], nums[i] = nums[i], nums[gt]
		}
	}

	nums[left], nums[lt] = nums[lt], nums[left]

	quick1(nums, left, lt-1)
	quick1(nums, lt+1, right)
}

func sort(nums []int) {
	tmp := make([]int, len(nums))
	mergeSort(nums, tmp, 0, len(nums)-1)
}

func mergeSort(nums, tmp []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	mergeSort(nums, tmp, left, mid)
	mergeSort(nums, tmp, mid+1, right)

	merge1(nums, tmp, left, mid, right)
}

func merge1(nums, tmp []int, left, mid, right int) {
	i := left
	j := mid + 1
	k := 0

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}

	for j <= right {
		tmp[k] = nums[j]
		j++
		k++
	}

	k = left
	for k <= right {
		nums[k] = tmp[k-left]
		k++
	}
}

func sort1(nums []int) {
	tmp := make([]int, len(nums))
	mergeSort1(nums, tmp, 0, len(nums)-1)
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
