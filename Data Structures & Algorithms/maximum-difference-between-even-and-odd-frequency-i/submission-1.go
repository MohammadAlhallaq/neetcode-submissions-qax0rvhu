func maxDifference(s string) int {

	freq := map[byte]int{}
	largestOdd := 0
	smallestEven := len(s)

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	for _, v := range freq {
		if v%2 == 1 {
			largestOdd = max(largestOdd, v)
		} else {
			smallestEven = min(smallestEven, v)
		}
	}
	return largestOdd - smallestEven
}
