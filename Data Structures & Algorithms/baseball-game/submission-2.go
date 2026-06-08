func calPoints(operations []string) int {
	stack := []int{}
	sum := 0

	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "+":
			first := stack[len(stack)-2]
			second := stack[len(stack)-1]
			val := first + second
			stack = append(stack, val)
			sum += val

		case "C":
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum -= last

		case "D":
			num := stack[len(stack)-1] * 2
			stack = append(stack, num)
			sum += num

		default:
			num, _ := strconv.Atoi(operations[i])
			stack = append(stack, num)
			sum += num
		}
	}
	return sum
}
