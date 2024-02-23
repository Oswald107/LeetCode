impl Solution {
    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        let n = matrix.len();
        for i in 0..n/2+n%2 {
            for j in 0..n/2 {
                let mut t1 = matrix[i][j];
                let mut a = i;
                let mut b = j;
                for _ in 0..4  {
                    let t2 = matrix[b][n-a-1];
                    matrix[b][n-a-1] = t1;
                    t1 = t2;
                    
                    let t3 = n-a-1;
                    a = b;
                    b = t3;
                }
            }
        }
    }
}