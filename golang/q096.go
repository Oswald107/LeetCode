func numTrees(n int) int {
	combos := make([]int, n+1)
	combos[0] = 1
	combos[1] = 1

	for i := 2; i <= n; i++ {
		v := 0
		for j := 1; j <= i; j++ {
			v += combos[j-1] * combos[i-j]
		}
		combos[i] = v
	}
	return combos[n]
}