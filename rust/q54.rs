impl Solution {
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        let n = matrix.len();
        let m = matrix[0].len();
        let mut dir = 0; // right-0, down-1, left-2, up-3
        let mut target = (0, m-1);
        let mut cur = (0, 0);
        let mut output = vec![];

        for _ in 0..n*m{
            output.push(matrix[cur.0][cur.1]);
            
            if cur == target {
                dir = (dir+1)%4;
                if dir == 0 {
                    target = (cur.0, m-2-cur.1)
                } else if dir == 2 {
                    target = (cur.0, m-1-cur.1);
                } else if dir == 1 {
                    target = (n-1-cur.0, cur.1);
                } else {
                    target = (n-cur.0, cur.1);
                }
            }

            if dir == 0 { cur.1 += 1; }
            else if dir == 1 { cur.0 += 1; }
            else if dir == 2 { cur.1 -= 1; }
            else { cur.0 -= 1; }
        }
        return output
    }
}