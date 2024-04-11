func longestValidParentheses(s string) int {
    var valid []Valid
    max := 0

    for i, v := range s {
        // we're looking for closed parenthesis to start the check
        if string(v) == "(" {
            continue
        }

        start := i
        isValid := false
        usedPrev := false

        // check if prev index has valid entry
        if len(valid) > 0 && valid[len(valid)-1].end == start-1 {
            start = valid[len(valid)-1].start
            usedPrev = true
        }

        // check if current substr is actually valid
        if start > 0 && string(s[start-1]) == "(" {
            start--
            isValid = true
            if usedPrev {
                valid = valid[:len(valid)-1]
            }
        }

        // check if prev index has valid entry
        if len(valid) > 0 && valid[len(valid)-1].end == start-1 {
            start = valid[len(valid)-1].start
            valid = valid[:len(valid)-1]
        }

        // append entry to valid list
        if isValid {
            valid = append(valid, Valid{start, i})
            if max < i-start+1 {
                max = i-start+1
            }
        }        
    }

    return max
}

type Valid struct {
    start int
    end int
}