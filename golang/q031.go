func nextPermutation(nums []int)  {
    var left, right int
    desc_ordered := true

    // get left value to swap
    for i := len(nums)-1; i > 0; i-- {
        if nums[i-1] < nums[i] {
            left = i-1
            desc_ordered = false
            break
        }
    }

    if desc_ordered {
        reverse(nums, 0, len(nums)-1)
        return
    }

    // get right value to swap
    right = left + 1
    for right < len(nums) - 1 {
        if nums[left] >= nums[right+1] {
            break
        }
        right++
    }

    swap(nums, left, right)
    reverse(nums, left+1, len(nums)-1)
}

func reverse(n []int, i int, j int)  {
    for i < j {
        swap(n, i, j)
        i++
        j--
    }
}

func swap(n []int, i int, j int) {
    temp := n[i]
    n[i] = n[j]
    n[j] = temp
}