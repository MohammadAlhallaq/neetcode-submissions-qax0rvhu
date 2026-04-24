
func dailyTemperatures(temperatures []int) []int {

	stack := make([]int, len(temperatures))
	n := len(temperatures)
	result := make([]int, n)

	for k := 0; k < n; k++ {
        for len(stack) > 0 && temperatures[k] > temperatures[stack[len(stack)-1]] {
				prevIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result[prevIndex] = k - prevIndex

		}
		stack = append(stack, k)
	}
	return result
}

// func dailyTemperatures(temperatures []int) []int {

// 	n := len(temperatures)
// 	res := make([]int, n)

// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			if temperatures[i] < temperatures[j] {
// 				res[i] = j - i
// 				break
// 			}
// 		}
// 	}
// 	return res
// }
