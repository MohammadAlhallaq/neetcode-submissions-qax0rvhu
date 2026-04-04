func longestConsecutive(nums []int) int {

	res := make(map[int]bool)
	longest := 0

	for _, v := range nums {
		res[v] = true
	}

	for v := range res {
		if _, ok := res[v - 1]; !ok {
			current := v
			length := 0


			for res[current] {
				length++
                current++
			}

			longest = max(longest, length)

		}
	}

	return longest
}
