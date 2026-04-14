func sortArrayByParity(nums []int) []int {
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right]%2 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
	return nums
}