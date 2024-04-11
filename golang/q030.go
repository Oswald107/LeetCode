func findSubstring(s string, words []string) []int {
    // word lengths
    word_len := len(words[0])
    total := len(words) * word_len

    // turn words into map for fast lookups
    m := make(map[string]int)
    for _, word := range words {
        _, ok := m[word]
        if ok {
            m[word]++
        } else {
            m[word] = 1
        }
    }

    // loop
    var output []int
    for i := 0; i < len(s)-total+1; i++ {
        if check(s, m, i, word_len, total) {
            output = append(output, i)
        }
    }
    return output
}


func check(s string, m map[string]int, i int, word_len int, total int) bool{
    m2 := make(map[string]int)
    for j := i; j < i + total; j+=word_len {
        sub := s[j:j+word_len]
        v1, ok := m[sub]
        v2, ok2 := m2[sub]
        if !ok || (ok2 && v1 == v2) {
            return false
        }

        if ok2 {
            m2[sub]++
        } else {
            m2[sub] = 1
        }
    }
    return true
}