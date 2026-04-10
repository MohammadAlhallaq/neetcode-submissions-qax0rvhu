func countSeniors(details []string) int {
	count := 0
	for i := 0; i < len(details); i++ {
		currentAge := ""
		j := 12
		for len(currentAge) < 2 {
			currentAge = string(details[i][j]) + currentAge
			j--
		}
		age, _ := strconv.Atoi(currentAge)
		if age > 60 {
			count++
		}
	}
	return count
}