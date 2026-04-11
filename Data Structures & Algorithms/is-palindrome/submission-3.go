func isPalindrome(s string) bool {
	k := len(s) - 1
	i := 0

	for i < k {
		for i < k && !isAlphaNumeric(rune(s[i])) {
			i++
		}

		for i < k && !isAlphaNumeric(rune(s[k])) {
			k--
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[k])) {
			return false
		}
		k--
		i++
	}
	return true
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}