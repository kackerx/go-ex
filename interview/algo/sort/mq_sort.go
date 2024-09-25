package sort

func Maopao(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		var flag bool
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				flag = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

		if !flag {
			break
		}
	}
}

// --------------------------------------------------------------------
func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2

	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)

	merge(arr, l, mid, r)
}

func merge(arr []int, L int, mid int, R int) {
	tmp := make([]int, 0, R-L+1)
	i := L
	j := mid + 1

	for i <= mid && j <= R {
		if arr[i] <= arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}

	for i <= mid {
		tmp = append(tmp, arr[i])
		i++
	}

	for j <= R {
		tmp = append(tmp, arr[j])
		j++
	}

	for k := L; k <= R; k++ {
		arr[k] = tmp[k-L]
	}
}

// --------------------------------------------------------------------
func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partitionV2(arr, l, r)
	QuickSort(arr, l, p-1)
	QuickSort(arr, p+1, r)
}

func partition(arr []int, L, R int) int {
	target := arr[L]
	j := L // <v = [L-1, j], >v = [j+1, i-1]

	for i := L + 1; i <= R; i++ {
		if arr[i] < target {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[L], arr[j] = arr[j], arr[L]

	return j
}

func partitionV2(arr []int, L, R int) int {
	target := arr[L]
	i := L + 1
	j := R

	for {
		for i <= j && arr[i] < target {
			i++
		}
		for i <= j && arr[j] > target {
			j--
		}

		// >: 遍历完
		// =: 等于的话说明这个值等于标定点因为i-->小于等于target, j-->大于等于target , i == j == target
		if i >= j {
			break
		}

		// 两层循环过后,
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	arr[L], arr[j] = arr[j], arr[L]
	return j
}

func partitionV3(arr []int, L, R int) {
	if L >= R {
		return
	}

	target := arr[L]
	i := L + 1
	lt := L
	gt := R + 1

	for i < gt {
		if arr[i] == target {
			i++
		} else if arr[i] < target {
			lt++
			arr[lt], arr[i] = arr[i], arr[lt]
			i++
		} else {
			gt--
			arr[gt], arr[i] = arr[i], arr[gt]
		}
	}

	arr[lt], arr[L] = arr[L], arr[lt]

	partitionV3(arr, L, lt-1)
	partitionV3(arr, gt, R)
}
