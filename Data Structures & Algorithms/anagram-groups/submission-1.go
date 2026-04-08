func groupAnagrams(strs []string) [][]string {


    m := make(map[[26]int][]string)

	for _, v := range strs {

		var freq [26]int

		for _, s := range v {
			freq[s-'a']++
		}	

		m[freq] = append(m[freq], v)
	}

	
	result := make([][]string, 0, len(m))

	for _, v := range m {
		result = append(result, v)
	}

	return result


}
