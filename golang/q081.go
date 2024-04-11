func search(nums []int, target int) bool {
	fmt.Println(nums)
	if len(nums) == 0 {
		return false
	}

	mid := len(nums) / 2
	if nums[mid] == target {
		return true
	}

	if nums[0] == nums[mid] && nums[mid] == nums[len(nums)-1] {
		return search(nums[:mid], target) || search(nums[mid+1:], target)
	}

	if nums[0] <= nums[mid] {
		if nums[0] <= target && target < nums[mid] {
			return search(nums[:mid], target)
		}
		return search(nums[mid+1:], target)
	} else {
		if nums[mid] < target && target <= nums[len(nums)-1] {
			return search(nums[mid+1:], target)
		}
		return search(nums[:mid], target)
	}
}