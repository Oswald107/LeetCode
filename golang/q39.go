func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	var output [][]int
	recur(&output, candidates, []int{}, target)
	return output
}

func recur(output *[][]int, nums []int, cur_collection []int, target int) {
	if target == 0 {
		cur_copy := make([]int, len(cur_collection))
		copy(cur_copy, cur_collection)
		*output = append(*output, cur_copy)
		return
	}
	for i, v := range nums {
		if target-v < 0 {
			return
		}
		recur(output, nums[i:], append(cur_collection, v), target-v)
	}
}