package sort

func InsertSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i; j-1 >= 0 && data[j-1] > data[j]; j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}

func InsertSortV2(data []int) {
	for i := 0; i < len(data); i++ {
		tmp := data[i]
		j := i
		for ; j-1 >= 0 && tmp < data[j-1]; j-- { // 只发生赋值, 最终复制给J位置暂存的tmp
			data[j] = data[j-1]
		}

		data[j] = tmp
	}
}
