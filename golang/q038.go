func countAndSay(n int) string {
    return recur("0", n)
}

func recur(num string, n int) string {
    if n == 0 {
        return num
    }

    if num == "0" {
        return recur("1", n-1)
    }

    output := ""
    cur := string(num[0])
    count := 0
    for _, v := range num {
        if string(v) == cur {
            count++
            continue
        }

        output += strconv.Itoa(count) + cur
        cur = string(v)
        count = 1
    }
    output += strconv.Itoa(count) + cur

    return  recur(output, n-1)
}