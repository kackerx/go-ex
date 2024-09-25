package search

func BinarySearch(arr []int, target int) int {
	l := 0
	r := len(arr)

	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func BinarySearchRecursion(arr []int, l, r, target int) int {
	if l > r {
		return -1
	}

	mid := l + (r-l)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return BinarySearchRecursion(arr, mid+1, r, target)
	} else {
		return BinarySearchRecursion(arr, l, mid-1, target)
	}

	return -1
}

// Upper 大于target的最小值
func Upper(arr []int, target int) int {
	l := 0
	r := len(arr)

	// [l, r]查找
	for l < r { // =的时候就找到目标了不需要继续循环了
		mid := l + (r-l)/2
		if arr[mid] <= target {
			l = mid + 1
		} else {
			r = mid // 此时mid有可能就是大于targe的最小值的可能是解, 不能刨除
		}
	}

	return l // 此时l==r, 返回谁一样
}

// Ceil 向上取整, 存在返回最大索引, 不存在返回upper
func Ceil(arr []int, target int) int {
	u := Upper(arr, target)

	if u-1 >= 0 && arr[u-1] == target {
		return u - 1
	}

	return u
}

// LowerCeil 大于等于target的最小索引
func LowerCeil(arr []int, target int) int {
	return -1
}

// Lower 小于target的最大值
func Lower(arr []int, target int) int {
	l := -1
	r := len(arr) - 1

	for l < r {
		mid := l + (r-l+1)/2   //
		if arr[mid] < target { // 小于的话这个mid可以不是小于target的最大, 所以不排除
			l = mid
		} else { // 对于>=的情况, 不符合, 排除 r = mid - 1
			r = mid - 1
		}
	}

	return l
}
