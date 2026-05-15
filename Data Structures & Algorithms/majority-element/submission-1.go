func majorityElement(nums []int) int {
	count := 0
	num := 0

	for _, v := range nums {
		if count == 0 {
			num = v
		}

		if num == v {
			count++
		} else {
			count--
		}
	}
	return num
}