package main

// leetcode submit region begin(Prohibit modification and deletion)
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := 0

	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
	}

	if sum < 0 { // 总加油不够总消耗, 肯定走不完
		return -1
	}

	tank := 0
	start := 0
	for i := 0; i < n; i++ {
		tank += gas[i] - cost[i]
		if tank < 0 { // 油量不够到当前为止, 那么就从下一个开始走
			tank = 0
			start = i + 1
		}
	}

	return start
}

// leetcode submit region end(Prohibit modification and deletion)
