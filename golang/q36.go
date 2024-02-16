func isValidSudoku(board [][]byte) bool {
    for i, row := range board {
        for j, val := range row  {
            if string(val) == "." {
                continue
            }
            for k := j+1; k < 9; k++ {
                if val == row[k] {
                    return false
                }
            }
            for k := i+1; k < 9; k++ {
                if val == board[k][j] {
                    return false
                }
            }
            for k := (i%3)*3 + (j%3) + 1; k < 9; k++ {
                if val == board[k/3 + (i/3 * 3)][k%3 + (j/3 * 3)] {
                    return false
                }
            }
        }
    }
    return true
}

