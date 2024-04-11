func generate(numRows int) [][]int {
	var output [][]int
	for i := 1; i <= numRows; i++ {
		var level []int

		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				level = append(level, 1)
			} else {
				level = append(level, output[len(output)-1][j-1]+output[len(output)-1][j])
			}
		}

		output = append(output, level)
	}
	return output
}