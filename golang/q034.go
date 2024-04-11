func searchRange(nums []int, target int) []int {
    s := start(nums, target)
    if s == -1 {
        return []int{-1, -1}
    }
    return []int{s, end(nums, target)}
}

func start(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    middleI := len(nums) / 2
    middle := nums[middleI]
    if nums[middleI] == target && (middleI == 0 || nums[middleI-1] != target){
        return middleI
    }

    if target <= middle {
        return start(nums[:middleI], target)
    } else {
        output := start(nums[middleI+1:], target)
        if output == -1 {
            return -1
        }
        return output + middleI + 1
    }
}

func end(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    middleI := len(nums) / 2
    middle := nums[middleI]
    if nums[middleI] == target && (len(nums)-1 == middleI || nums[middleI+1] != target){
        return middleI
    }

    if target < middle {
        return end(nums[:middleI], target)
    } else {
        output := end(nums[middleI+1:], target)
        if output == -1 {
            return -1
        }
        return output + middleI + 1
    }
}