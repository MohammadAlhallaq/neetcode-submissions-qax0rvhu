type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	for i := 0; i < len(strs); i++ {
		curr := strs[i]
		length := len(curr)
		encoded += strconv.Itoa(length) + "#" + curr
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	index := 0
	res := []string{}

	// 4#neet4#code4#love3#you
	for index < len(encoded) {
		j := index
		for j < len(encoded) && unicode.IsDigit(rune(encoded[j])) {
			j++
		}
		numserStr := encoded[index:j]
		number, _ := strconv.Atoi(numserStr)
		start := j + 1
		end := start + number
		res = append(res, encoded[start:end])
		index = end
	}
	return res
}