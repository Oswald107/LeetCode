func removeElement(nums []int, val int) int {
    k := 0
    for _, v := range nums {
        if val != v {
            nums[k] = v
            k += 1
        }
    }
    return k
}