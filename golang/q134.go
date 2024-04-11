func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	start := l - 1
	end := l - 1
	cur := gas[end] - cost[end]
	for start != (end+1)%l {
		if cur <= 0 {
			start--
			cur += gas[start] - cost[start]
		} else {
			end = (end + 1) % l
			cur += gas[end] - cost[end]
		}
	}

	if cur >= 0 {
		return start
	}
	return -1
}