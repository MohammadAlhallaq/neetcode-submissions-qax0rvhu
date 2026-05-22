func minimumRecolors(blocks string, k int) int {
	whiteCount := 0

	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			whiteCount++
		}
	}

	minOps := whiteCount

	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			whiteCount++
		}
		if blocks[i-k] == 'W' {
			whiteCount--
		}

		minOps = min(minOps, whiteCount)
	}

	return minOps
}
