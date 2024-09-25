package sort

func BubbleSort(data []int) {
	for i := 0; i < len(data)-1; i++ { // 不需要全部遍历n遍, 因为最后一次只剩一个元素就不用比较了
		for j := 0; j+1 <= len(data)-1-i; j++ { // 这里终止条件理解: 我们每轮比较的是j和j+1, 因此拿到j+1就行, j+1 <= len-1, 去掉每轮的i

		}

	}
}
