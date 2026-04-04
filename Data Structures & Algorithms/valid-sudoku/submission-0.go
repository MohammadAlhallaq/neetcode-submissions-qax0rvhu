func isValidSudoku(board [][]byte) bool {

	res := make(map[string]bool)

	for i := 0; i < 9; i++{
		for j := 0; j < 9; j++{
			curr := board[i][j]
			if curr != '.'{

				rowKey := string(curr) + "row" + strconv.Itoa(i)
				colKey := string(curr) + "col" + strconv.Itoa(j)
				boxKey := string(curr) + "box" + strconv.Itoa(i/3) + "-" + strconv.Itoa(j/3)

				if !addToSeen(rowKey, res) ||
					!addToSeen(colKey, res) ||
					!addToSeen(boxKey, res) {
					return false
				}
			}
		}
	}
	return true
}

func addToSeen(element string, res map[string]bool) bool{

	if _, ok := res[element]; ok {
		return false // already exists → invalid
	}

	res[element] = true
	return true
}
