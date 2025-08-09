package main_test

func isValidSudoku(board [][]byte) bool {
	for _, row := range board {
		m := make(map[string]struct{})
		for _, cell := range row {
			s := string(cell)
			if s == "." {
				continue
			}
			if _, ok := m[s]; ok {
				return false
			} else {
				m[s] = struct{}{}
			}
		}
	}
	for i := range len(board) {
		m := make(map[string]struct{})
		for j := range len(board) {
			s := string(board[j][i])
			if s == "." {
				continue
			}
			if _, ok := m[s]; ok{
				return false
			} else {
				m[s] = struct{}{}
			}
		}
	}
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board); j += 3 {
			m := make(map[string]struct{})
			for k := range 3 {
				for l := range 3 {
					s := string(board[i+k][j+l])
					if s == "." {
						continue
					}
					if _, ok := m[s]; ok {
						return false
					} else {
						m[s] = struct{}{}
					}

				}
			}
		}
	}
	return true
}
