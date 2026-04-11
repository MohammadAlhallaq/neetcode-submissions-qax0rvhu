func mergeAlternately(word1 string, word2 string) string {
	i, j := 0, 0
	n, m := len(word1), len(word2)

	res := ""

	for i < n || j < m {
		if i < n {
			res += string(word1[i])
			i++
		}
		if j < m {
			res += string(word2[j])
			j++
		}
	}

	return res
}