impl Solution {
    pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut output = Vec::new();
        recur(nums, vec![], &mut output);
        return output;
    }
}

fn recur(nums: Vec<i32>, path: Vec<i32>, output: &mut Vec<Vec<i32>>) {
    if nums.len() == 0{
        output.push(path);
        return;
    }

    for i in 0..nums.len() {
        let mut nums_clone = nums.clone();
        let mut path_clone = path.clone();
        nums_clone.remove(i);
        path_clone.push(nums[i]);
        recur(nums_clone, path_clone, output);
    }
}