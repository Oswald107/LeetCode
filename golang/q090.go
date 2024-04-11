func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var output [][]int
	recur(&output, []int{}, nums, 0)
	return output
}

func recur(output *[][]int, cur []int, nums []int, i int) {
	// append cur to output, use a copy to prevent overwriting
	dst := make([]int, len(cur))
	copy(dst, cur)
	*output = append(*output, dst)

	for i < len(nums) {
		v := nums[i]
		// count time value v appears
		count := 1
		for i+1 < len(nums) && nums[i+1] == v {
			i++
			count++
		}
		// recur through each, but the starting point should be the next new value
		j := 0
		for count > j {
			cur = append(cur, v)
			recur(output, cur, nums, i+1)
			j++
		}
		// remove all copys of v
		cur = cur[:len(cur)-count]
		i++
	}
}