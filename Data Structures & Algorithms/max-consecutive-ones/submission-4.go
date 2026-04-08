func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	top := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			top = max(count, top)
		} else {
			count = 0
		}
	}
	return top
}