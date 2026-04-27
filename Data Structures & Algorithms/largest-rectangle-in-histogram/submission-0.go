func largestRectangleArea(heights []int) int {
	stack := []int{} // stores indices
	maxArea := 0
	// 0: ███████ 7
	// 1: █ 1
	// 2: ███████ 7
	// 3: ██ 2
	// 4: ██ 2
	// 5: ████ 4

	// [7]
    // i = 1
	// add a sentinel (0 height) to flush the stack at the end
	heights = append(heights, 0)

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := heights[top]

			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}

	return maxArea
}