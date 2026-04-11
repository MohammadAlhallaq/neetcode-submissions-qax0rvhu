func validPalindrome(s string) bool {
	k := len(s) - 1
	i := 0

	for i < k {
		for i < k && !unicode.IsLetter(rune(s[i])) {
			i++
		}

		for i < k && !unicode.IsLetter(rune(s[k])) {
			k--
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[k])) {
			return isPalindromeRange(s, i+1, k) || isPalindromeRange(s, i, k-1)
		}
		k--
		i++
	}
	return true
}

func isPalindromeRange(s string, i, k int) bool {
	for i < k {
		for i < k && !unicode.IsLetter(rune(s[i])) {
			i++
		}

		for i < k && !unicode.IsLetter(rune(s[k])) {
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