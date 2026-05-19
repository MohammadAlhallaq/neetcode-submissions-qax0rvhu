func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	freq := map[byte]bool{}
	slow := 0

	for fast := 0; fast < len(s); fast++ {
		for freq[s[fast]] {
			delete(freq, s[slow])
			slow++
		}

		freq[s[fast]] = true

		maxLength = max(maxLength, fast-slow+1)
	}
	return maxLength
}
