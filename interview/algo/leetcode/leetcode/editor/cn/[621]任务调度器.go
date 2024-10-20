package main

// leetcode submit region begin(Prohibit modification and deletion)
func leastInterval(tasks []byte, n int) int {
	// 如图 --> 总排队时间 = (桶个数 - 1) * (n + 1) + 最后一桶的任务数
	ht := [26]int{}
	N := 0     // 最多任务字母
	count := 0 // 最多任务字母的个数

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


// leetcode submit region end(Prohibit modification and deletion)
