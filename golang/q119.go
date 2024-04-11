func getRow(rowIndex int) []int {
	prev := []int{1}
	if rowIndex == 0 {
		return prev
	}

	var level []int
	for i := 1; i <= rowIndex; i++ {
		level = []int{}

		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				level = append(level, 1)
			} else {
				level = append(level, prev[j-1]+prev[j])
			}
		}

		prev = level
	}

	return level
}