impl Solution {
    pub fn permute_unique(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut output = Vec::new();
        let mut nums_clone = nums;
        nums_clone.sort();
        recur(nums_clone, vec![], &mut output);
        return output;
    }
}

fn recur(nums: Vec<i32>, path: Vec<i32>, output: &mut Vec<Vec<i32>>) {
    if nums.len() == 0{
        output.push(path);
        return;
    }

    let mut prev = -11;
    for i in 0..nums.len() {
        if nums[i] == prev {
            continue
        }
        let mut nums_clone = nums.clone();
        let mut path_clone = path.clone();
        prev = nums_clone.remove(i);
        path_clone.push(prev);
        recur(nums_clone, path_clone, output);
    }
}