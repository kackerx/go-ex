package main

// leetcode submit region begin(Prohibit modification and deletion)
func corpFlightBookings(bookings [][]int, n int) []int {
	d := NewDifference(make([]int, n))

	for _, booking := range bookings {
		d.Increment(booking[0]-1, booking[1]-1, booking[2])
	}

	return d.Res()
}

type Difference struct {
	diff []int
}

func NewDifference(nums []int) *Difference {
	diff := make([]int, len(nums))

	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{diff: diff}
}

func (d *Difference) Increment(i, j, inc int) {
	d.diff[i] += inc // 这一步使得nums[i:]全部加3, 也包括了j后面的
	if j+1 < len(d.diff) {
		d.diff[j+1] -= inc // 这一步把j后面的全部减去3, 就实现了只[i, j]范围+3
	}
}

func (d *Difference) Res() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]

	for i := 1; i < len(d.diff); i++ {
		res[i] = d.diff[i] + res[i-1]
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
