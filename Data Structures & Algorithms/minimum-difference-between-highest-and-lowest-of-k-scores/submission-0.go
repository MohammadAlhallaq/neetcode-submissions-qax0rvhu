func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	res := math.MaxInt
	// [2,5,3,1,6,3]
	for i := k - 1; i < len(nums); i++ {
		res = min(res, nums[i]-nums[i-k+1])
	}

	return res
}