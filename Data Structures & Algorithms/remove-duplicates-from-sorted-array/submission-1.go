func removeDuplicates(nums []int) int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}

	return k+1
}