func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {

		temp := numbers[left] + numbers[right]

		if temp == target {
			return []int{left + 1, right + 1}
		}

		if temp > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}