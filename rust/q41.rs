impl Solution {
    pub fn first_missing_positive(mut nums: Vec<i32>) -> i32 {
        let mut i: usize = 0;
        while i < nums.len() {
            if i+1 != nums[i] as usize && nums[i] as usize <= nums.len() && nums[i] > 0{
                let temp = nums[i];
                if temp != nums[(temp-1) as usize] {
                    nums[i] = nums[(temp-1) as usize];
                    nums[(temp-1) as usize] = temp;
                } else {
                    i+=1;
                }
                continue;
            }
            i+=1;
        }

        for i in 0..nums.len() {
            if i+1 != nums[i] as usize{
                return (i+1) as i32;
            }
        }
        return (nums.len() + 1) as i32;
    }
}