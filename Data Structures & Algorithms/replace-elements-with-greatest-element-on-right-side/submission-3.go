func replaceElements(arr []int) []int {
	maxRight := -1

	for i := len(arr) - 1; i >= 0; i-- {
		curr := arr[i]
		arr[i] = maxRight
		maxRight = max(maxRight, curr)
	}
	
	return arr
}