func longestMonotonicSubarray(nums []int) int {

	maxDec := 1
	maxInc := 1
	currInc := 1
	currDec := 1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			currInc++
			currDec = 1
		} else if nums[i] > nums[i+1] {
			currDec++
			currInc = 1
		} else {
			currInc = 1
			currDec = 1
		}
		maxDec = max(maxDec, currDec)
		maxInc = max(maxInc, currInc)

	}
	return max(maxDec, maxInc)
}
