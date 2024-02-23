impl Solution {
    pub fn generate_matrix(n: i32) -> Vec<Vec<i32>> {
        let mut output = vec![];
        let n = n as usize;
        for _ in 0..n {
            output.push(vec![0; n]);
        }

        let mut dir = 0; // right-0, down-1, left-2, up-3
        let mut target = (0, n-1);
        let mut cur = (0, 0);

        for i in 1..n*n+1{
            output[cur.0][cur.1] = i as i32;
            
            if cur == target {
                dir = (dir+1)%4;
                if dir == 0 {
                    target = (cur.0, n-2-cur.1)
                } else if dir == 2 {
                    target = (cur.0, n-1-cur.1);
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
        return output;
    }
}
