impl Solution {
    pub fn merge(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut intervals = intervals.clone();
        let mut output = vec![];
        intervals.sort();
        
        let mut cur = intervals[0].clone();
        for i in 1..intervals.len() {
            let next = intervals[i].clone();
            if cur[1] < next[0] {
                output.push(cur);
                cur = intervals[i].clone();
            } else if cur[1] < next[1] {
                cur[1] = next[1];
            }
        }
        output.push(cur);

        return output;
    }
}