func majorityElement(nums []int) int {
	n := len(nums)
	freq := map[int]int{}
	half := int(math.Floor(float64(n / 2)))

	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	for k, v := range freq {
		if v > half {
			return k
		}
	}
	return 0
}