impl Solution {
    pub fn jump(nums: Vec<i32>) -> i32 {
        let mut output = 1;
        let mut i = 0;
        let mut j = nums[0];
        
        if nums.len() as i32 == 1 {
            return 0
        }

        while i+j < nums.len() as i32 - 1  {
            let mut max = 0;
            let mut max_index = 0;
            
            for k in i+1..i+j+1 {
                if nums[k as usize] > max - k + max_index {
                    max = nums[k as usize];
                    max_index = k;
                }
            }

            j = max;
            i = max_index;
            output+=1;
        }
        
        return output
    }
}