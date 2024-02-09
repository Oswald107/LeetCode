func removeDuplicates(nums []int) int {
    var prev int
    output := 0

    for i, v := range nums {
        if i ==0  || prev != v {
            nums[output] = v
            output += 1
        }
        prev = v
    }
    return output
}