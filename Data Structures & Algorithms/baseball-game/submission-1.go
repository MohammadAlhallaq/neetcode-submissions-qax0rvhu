func calPoints(operations []string) int {
	stack := []int{}
	sum := 0

	for i := 0; i < len(operations); i++ {
		if len(stack) == 0 {
			num, _ := strconv.Atoi(operations[i])
			stack = append(stack, num)
			continue
		}

		if operations[i] == "+" {
			first := stack[len(stack)-2]
			second := stack[len(stack)-1]
			stack = append(stack, first+second)
			fmt.Print(stack)
		} else if operations[i] == "C" {
			stack = stack[:len(stack)-1]
			fmt.Print(stack)

		} else if operations[i] == "D" {
			num := stack[len(stack)-1]
			stack = append(stack, num*2)
			fmt.Print(stack)
		} else {
			num, _ := strconv.Atoi(operations[i])
			stack = append(stack, num)
			fmt.Print(stack)
		}
	}

	for i := 0; i < len(stack); i++ {
		sum += stack[i]
	}

	return sum
}
