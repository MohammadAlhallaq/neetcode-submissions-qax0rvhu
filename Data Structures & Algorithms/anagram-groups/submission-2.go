func groupAnagrams(strs []string) [][]string {

	m := make(map[[26]int][]string)

	for i := 0; i < len(strs); i++ {
		curr := strs[i]
		var freq [26]int

		for _, v := range curr {
			freq[v-'a']++
		}
		m[freq] = append(m[freq], curr)
	}

	res := make([][]string, 0, len(m))

	for _, v := range m {
		res = append(res, v)
	}
	return res
}