package array

import (
	"fmt"
	"testing"
)

func Test621(t *testing.T) {
	s := []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C'}
	fmt.Println(leastInterval(s, 2))
}

func leastInterval(tasks []byte, n int) int {
	// 如图 --> 总排队时间 = (桶个数 - 1) * (n + 1) + 最后一桶的任务数
	ht := [26]int{}
	N := 0     // 最多任务字母
	count := 0 // 最多任务字母的种类数

	for _, task := range tasks {
		ht[task-'A']++
	}

	for i := 0; i < 26; i++ {
		if ht[i] > N {
			N = ht[i]
			count = 1
		} else if ht[i] == N {
			count++
		}
	}

	return max(len(tasks), (n+1)*(N-1)+count)
}
