func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))


	result[0] = 1
	for i := 1; i < len(nums); i++{
		result[i] = result[i - 1] * nums[i - 1]

	}


	suffix := 1
	for i := len(nums) - 1; i >= 0; i--{
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

// func productExceptSelf(nums []int) []int {
// 	res := []int{}

// 	for i := 0; i < len(nums); i++ {
// 		temp := 1

// 		for j := 0; j < len(nums); j++ {
// 			if i != j {
// 				temp *= nums[j]
// 			}
// 		}

// 		res = append(res, temp)
// 	}

// 	return res
// }
