func trap(height []int) int {
	left := 0
	right := len(height) - 1
	maxL := height[left]
	maxR := height[right]
	res := 0

	for left < right {
		if maxL < maxR {
			left++
			maxL = max(maxL, height[left])
			res += maxL - height[left]
		} else {
			right--
			maxR = max(maxR, height[right])
			res += maxR - height[right]
		}
	}
	return res
}
