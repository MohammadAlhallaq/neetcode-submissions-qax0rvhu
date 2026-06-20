func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	leftProduct := 1
	for i := 0; i < n; i++ {
		res[i] = leftProduct
		leftProduct *= nums[i]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return res
}
