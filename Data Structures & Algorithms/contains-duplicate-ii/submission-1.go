func containsNearbyDuplicate(nums []int, k int) bool {
	res := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if _, exists := res[nums[i]]; exists {
			return true
		}

		res[nums[i]] = struct{}{}

		if len(res) > k {
			delete(res, nums[i-k])

		}
	}
	return false
}