func minOperations(logs []string) int {
	stack := []string{}

	for i := 0; i < len(logs); i++ {
		char := logs[i]
		if char == "../" && len(stack) != 0 {
			stack = stack[:len(stack)-1]
			continue
		} else if char == "./" || (char == "../" && len(stack) == 0) {
			continue
		}
		stack = append(stack, char)
	}
	return len(stack)
}