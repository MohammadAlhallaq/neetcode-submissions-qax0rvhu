func findContentChildren(g []int, s []int) int {
	i, j, res := 0, 0, 0
    sort.Ints(s)
    sort.Ints(g)

	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
			j++
			res++
		} else {
			j++
		}
	}

	return res
}