package sort

func SelectSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}

		data[i], data[minIndex] = data[minIndex], data[i]
	}
}
