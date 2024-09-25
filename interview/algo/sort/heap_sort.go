package sort

func HeapSort(data []int) {
	if len(data) <= 1 {
		return
	}

	for i := (len(data) - 2) / 2; i >= 0; i-- {
		siftDown(data, i, len(data))
	}

	for i := len(data) - 1; i >= 0; i-- {
		data[0], data[i] = data[i], data[0]
		siftDown(data, 0, i)
	}
}

// siftDown 对堆data[0, n)的k位的元素进行下沉
func siftDown(data []int, k, n int) {
	for 2*k+1 < n {
		j := 2*k + 1
		if j+1 < n && data[j+1] > data[j] {
			j++
		}

		if data[k] >= data[j] {
			break
		}

		data[k], data[j] = data[j], data[k]
		k = j
	}
}
