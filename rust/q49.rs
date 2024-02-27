use std::collections::HashMap;
impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {

        let mut map = HashMap::new();
        
        for s in strs {
            let mut chars: Vec<char> = s.chars().collect();
            chars.sort_unstable();
            let sorted_str: String = chars.into_iter().collect();

            map.entry(String::from(sorted_str)).or_insert(Vec::new()).push(s);
        }

        let mut output = vec![];
        for lst in map.values() {
            output.push(lst.to_vec());
        }
        return output;
    }
}