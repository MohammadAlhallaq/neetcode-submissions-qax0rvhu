func generate(numRows int) [][]int {

	res := [][]int{}

	for i := 0; i < numRows; i++ {
		row := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				num := res[i-1][j-1] + res[i-1][j]
				row = append(row, num)
			}
		}
		res = append(res, row)
	}
	return res
}
