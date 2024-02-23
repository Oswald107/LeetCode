impl Solution {
    pub fn get_permutation(n: i32, k: i32) -> String {
        return recur(n, k-1, (1..=n).collect())
    }
}

fn recur(n: i32, k: i32, nums: Vec<i32>) -> String {
    let mut output = "".to_string();
    if k == 0  || n == 1{
        for val in nums {
            output = format!("{}{}", output, val);
        }
        return output
    }

    let perm: i32 = (1..n).product();
    let i = k / perm;
    let k = k % perm;
    let mut nums = nums.clone();
    let val = nums.remove(i as usize);
    output = format!("{}{}{}", output, val, recur(n-1, k, nums));
    return output;
}