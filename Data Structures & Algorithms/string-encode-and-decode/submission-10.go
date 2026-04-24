type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	res := ""

	for i := 0; i < len(strs); i++ {
		res += strconv.Itoa(len(strs[i])) + "#" + strs[i]
	}

	return res
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0

	for i < len(encoded) {

		j := i
		for j < len(encoded) && unicode.IsDigit(rune(encoded[j])) {
			j++
		}

		numberStr := encoded[i:j]
		number, _ := strconv.Atoi(numberStr)
		start := j + 1
		end := start + number
		res = append(res, encoded[start:end])
		i = end
	}
	return res
}