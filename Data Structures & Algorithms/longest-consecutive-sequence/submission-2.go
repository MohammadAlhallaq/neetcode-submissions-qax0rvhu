func longestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums))
	longest := 0

	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		set[curr] = true
	}

	for v := range set {
		if _, ok := set[v-1]; !ok {
			current := v
			lenght := 0

			for set[current] {
				lenght++
				current++
			}

			longest = max(lenght, longest)
		}

	}
	return longest
}
