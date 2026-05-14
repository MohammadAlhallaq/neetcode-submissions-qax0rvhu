func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	res := 0

	for i := 0; i < len(heights); i++ {
		temp := (right - left) * min(heights[left], heights[right])
		res = max(temp, res)
		if heights[left] > heights[right] {
			right--
		} else {
			left++
		}
	}
	return res
}