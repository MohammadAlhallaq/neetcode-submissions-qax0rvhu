func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	n := len(nums)

	res := [][]int{}

	for i := 0; i < n; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			num := nums[left] + nums[i] + nums[right]
			if num == 0 {
				res = append(res, []int{
					nums[i],
					nums[left],
					nums[right],
				})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if num > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

