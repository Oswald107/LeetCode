impl Solution {
    pub fn insert(intervals: Vec<Vec<i32>>, new_interval: Vec<i32>) -> Vec<Vec<i32>> {
        let mut new_interval = new_interval.clone();
        let mut output = vec![];
        
        let mut i = 0;
        while i < intervals.len() {
            let cur = intervals[i].clone();
            i+=1;
            if cur[1] < new_interval[0] {
                output.push(cur);
                continue
            } if new_interval[1] < cur[0] {
                i-=1;
                break
            }
            
            if new_interval[1] < cur[1] {
                new_interval[1] = cur[1];
            }
            if cur[0] < new_interval[0] {
                new_interval[0] = cur[0];
            }
        }

        output.push(new_interval);

        for i in i..intervals.len() {
            let cur = intervals[i].clone();
            output.push(cur);
        }

        return output;
    }
}
