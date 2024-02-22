impl Solution {
    pub fn length_of_last_word(s: String) -> i32 {
        let char_bytes = s.as_bytes();
        let mut found = false;
        let mut end = s.len()-1;

        for i in (0..s.len()).rev() {
            if char_bytes[i] as char == ' ' {
                if found {
                    return (end-i) as i32;
                }
            } else if !found {
                found = true;
                end = i
            }
            
        }

        if char_bytes[0] as char == ' ' {
            return 0;
        }
        return (end+1) as i32;
    }
}