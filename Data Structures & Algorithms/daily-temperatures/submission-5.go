
func dailyTemperatures(temperatures []int) []int {

	stack := make([]int, len(temperatures))
	n := len(temperatures)
	result := make([]int, n)

	for k := 0; k < len(temperatures); k++ {
        for len(stack) > 0 && temperatures[k] > temperatures[stack[len(stack)-1]] {
				prevIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result[prevIndex] = k - prevIndex

		}
		stack = append(stack, k)
	}
	return result
}
