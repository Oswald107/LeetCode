
impl Solution {
    pub fn solve_n_queens(n: i32) -> Vec<Vec<String>> {
        let mut output = Vec::new();
        recur(n, 0, vec![], &mut output);
        return output;
    }
}


fn recur(n: i32, row: i32, placed: Vec<i32>, output: &mut Vec<Vec<String>>) {
    if n == row {
        output.push(make_board(placed));
        return;
    }
    
    for col in 0..n {
        if valid_space(row, col, placed.clone()) {
            let mut placed_clone = placed.clone();
            placed_clone.push(col);
            recur(n, row+1, placed_clone, output);
        }
    }
}

fn valid_space(row: i32, col: i32, placed: Vec<i32>) -> bool{
    let mut row2 = 0;
    for col2 in placed {
        if col == col2 {
            return false;
        }
        if (col2 - col).abs() == (row2 - row).abs() {
            return false
        }
        row2+=1
    }
    return true
}

fn make_board(placed: Vec<i32>) -> Vec<String> {
    let mut output = vec![];
    let l = placed.len();
    for p in placed {
        let mut s = "".to_string();
        for i in 0..l {
            if p == i as i32 {
                s.push_str("Q");
            } else {
                s.push_str(".");
            }
        }
        output.push(s);
    }
    return output;
}