func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	currMax := 0

	for left < right {
		width := right - left
		minHeight := min(heights[left], heights[right])
		temp := width * minHeight

		if temp > currMax {
			currMax = temp
		}

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return currMax
}