func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }
    return recur(nums, target, true)
}

func recur(nums []int, target int, forward bool) int {
    if len(nums) == 0 {
        return -1
    }

    middleI := len(nums) / 2
    middle := nums[middleI]
    if middle == target {
        return middleI
    }

    if target < middle { 
        output := recur(nums[:middleI], target, false)
        if output == -1 {
            return middleI
        }
        return output
    } else {
        output := recur(nums[middleI+1:], target, true)
        if output == -1 {
            return middleI+1
        }
        return output + middleI + 1
    }
}