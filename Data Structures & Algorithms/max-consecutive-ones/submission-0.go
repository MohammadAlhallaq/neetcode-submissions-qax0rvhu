func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	top := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if count > top {
				top = count
			}
		} else {
			count = 0
		}
	}
	return top
}