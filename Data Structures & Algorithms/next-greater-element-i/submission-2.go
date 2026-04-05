func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	nextGreater := make(map[int]int)

	// Build next greater map using nums2
	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreater[top] = num
		}
		stack = append(stack, num)
	}

	for _, num := range stack {
		nextGreater[num] = -1
	}


	result := make([]int, len(nums1))
	for k, v := range nums1{
		result[k] = nextGreater[v]
	}
	return result
}