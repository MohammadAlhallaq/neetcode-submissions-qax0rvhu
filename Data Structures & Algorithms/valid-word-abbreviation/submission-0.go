func validWordAbbreviation(word string, abbr string) bool {
	k := 0
	i := 0

	n := len(word)
	m := len(abbr)

	for i < n && k < m {
		if unicode.IsLetter(rune(abbr[k])) {
			if word[i] != abbr[k] {
				return false
			}
			i++
			k++
			continue
		}
		if abbr[k] == '0' {
			return false
		}

		num := 0
		for k < m && unicode.IsDigit(rune(abbr[k])) {
			num = num*10 + int(abbr[k]-'0')
			k++
		}

		i += num
	}

	return i == n && k == m
}