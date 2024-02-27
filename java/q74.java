class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        int m = matrix.length;
        int n = matrix[0].length;
        int start = 0;
        int end = m*n-1;

        while (start <= end) {
            int ind = (start + end) / 2;
            int i = ind/n;
            int j = ind%n;

            if (matrix[i][j] == target)
                return true;
            if (matrix[i][j] > target)
                end = ind-1;
            else
                start = ind + 1;
        }

        return false;
    }
}
