func search(nums []int, target int) int {
    return recur(nums, target)   
}

func recur(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    middleI := len(nums) / 2
    middle := nums[middleI]
    if nums[middleI] == target {
        return middleI
    }

    first := nums[0]
    last := nums[len(nums)-1]

    if first < middle { 
        if first <= target && target < middle {
            return recur(nums[:middleI], target)
        }
        output := recur(nums[middleI+1:], target)
        if output == -1 {
            return -1
        }
        return output + middleI + 1
    } else {
        if middle < target && target <= last {
            output := recur(nums[middleI+1:], target)
            if output == -1 {
                return -1
            }
            return output + middleI + 1
        }
        return recur(nums[:middleI], target)
    }
}