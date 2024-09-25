package leetcode

import (
	"fmt"

	"go-bobo/interview/algo/struct/queue"
)

func GetLeastNumbers(data []int, k int) {
	pq := queue.NewPriorityQueue()

	for i := 0; i < k; i++ { // 数组前K先入最大堆
		pq.Enqueue(data[i])
	}

	for i := k; i < len(data); i++ {
		if data[i] < pq.GetFront() { // 从k开始, 如果值比最大堆的最大还小, 那就替换掉最大, 堆维护最小的K个
			pq.Dequeue()
			pq.Enqueue(data[i])
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = pq.Dequeue()
	}

	fmt.Println(res)
}
