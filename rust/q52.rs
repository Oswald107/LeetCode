impl Solution {
    pub fn total_n_queens(n: i32) -> i32 {
        return recur(n, 0, vec![]);
    }
}


fn recur(n: i32, row: i32, placed: Vec<i32>) -> i32 {
    if n == row {
        return 1;
    }
    
    let mut output = 0;
    for col in 0..n {
        if valid_space(row, col, placed.clone()) {
            let mut placed_clone = placed.clone();
            placed_clone.push(col);
            output += recur(n, row+1, placed_clone);
        }
    }
    return output;
}

fn valid_space(row: i32, col: i32, placed: Vec<i32>) -> bool{
    let mut row2 = 0;
    for col2 in placed {
        if col == col2 {
            return false;
        }
        if (col2 - col).abs() == (row2 - row).abs() {
            return false;
        }
        row2+=1;
    }
    return true;
}