func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	res := []int{}
	for left < right {
		currSum := numbers[left] + numbers[right]
		if currSum == target {
			res = []int{left + 1, right + 1}
		}
		if currSum > target {
			right--
		} else {
			left++
		}
	}
	return res
}