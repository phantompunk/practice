package kata

// ::KATA START::
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := range 9 {
		rows[i] = make(map[byte]bool, 9)
		cols[i] = make(map[byte]bool, 9)
		boxes[i] = make(map[byte]bool, 9)
	}

	for r := range 9 {
		for c := range 9 {
			value := board[r][c]
			if value == byte('.') {
				continue
			}

			if rows[r][value] {
				return false
			}

			if cols[c][value] {
				return false
			}

			if boxes[(r/3)*3+c/3][value] {
				return false
			}

			rows[r][value] = true
			cols[c][value] = true
			boxes[(r/3)*3+c/3][value] = true
		}
	}
	return true
}

// ::KATA END::
