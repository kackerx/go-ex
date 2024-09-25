package leetcode

func SortColors(arr []int) {
	lt := -1
	gt := len(arr)
	i := 0
	for i < gt {
		if arr[i] == 1 {
			i++
		} else if arr[i] < 1 {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
		} else {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		}
	}
}
