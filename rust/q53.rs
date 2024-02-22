impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let mut max = std::i32::MIN;
        let mut sub = vec![];
        let mut total = 0;
        for v in nums.iter() {
            if total < 0 {
                total = 0;
                sub = vec![];
            }
            sub.push(v);
            total += v;
            if total > max {
                max = total;
            }
        }
        return max;
    }
}