func sortedSquares(nums []int) []int {
	n := len(nums)

	i, j := 0, n-1
	end := n - 1
	result := make([]int, n)

	for i <= j {
		sVal := nums[i] * nums[i]
		eVal := nums[j] * nums[j]
		if sVal > eVal {
			result[end] = sVal
			i++
		} else {
			result[end] = eVal
			j--
		}
		end--
	}
	return result
}
