func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = 0
	}

	max := 0
	for start, _ := range m {
		count := 1
		end := start

		// check increments, if val is > 0, we checked it already
		for true {
			end_v, ok := m[end+1]
			if !ok {
				break
			}
			if end_v > 0 {
				count += m[end+1]
				break
			}
			end++
			count++
		}

		// update all values in the chain
		for end >= start {
			m[end] = count - (end - start)
			end--
		}

		// fmt.Println(start, count)
		// update max
		if count > max {
			max = count
		}
	}
	return max
}