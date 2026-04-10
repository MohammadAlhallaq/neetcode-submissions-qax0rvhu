func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	res := []int{}
	hash := map[int]bool{}

	for i := 0; i < n; i++ {
		hash[nums[i]] = true
	}

	for i := 1; i <= n; i++ {
		if _, exists := hash[i]; !exists {
			res = append(res, i)
		}
	}
	return res
}
