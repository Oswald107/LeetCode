class Solution {
    public boolean exist(char[][] board, String word) {
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if (recur(board, word, new boolean[board.length][board[0].length], i, j, 0))
                    return true;
            }
        }
        return false;
    }

    public boolean recur(char[][] board, String word, boolean[][] explored, int i, int j, int c) {
        if (c == word.length())
            return true;
        if (i < 0 || j < 0 || i >= board.length || j >= board[0].length)
            return false;
        if (board[i][j] != word.charAt(c))
            return false;
        if (explored[i][j])
            return false;


        explored[i][j] = true;
        boolean output = recur(board, word, explored, i+1, j, c+1)
            || recur(board, word, explored, i-1, j, c+1)
            || recur(board, word, explored, i, j+1, c+1)
            || recur(board, word, explored, i, j-1, c+1);
        explored[i][j] = false;
        return output;
    }
}