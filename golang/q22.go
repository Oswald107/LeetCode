func generateParenthesis(n int) []string {
    var output []string
    recur(&output, n, n, "")
    return output

}

func recur(output *[]string, open int, close int, cur string) {
    if open == 0 && close == 0 {
        *output = append(*output, cur)
        return
    }

    if open < close {
        recur(output, open, close-1, cur+")")
    }

    if open > 0{
        recur(output, open-1, close, cur+"(")
    }
}