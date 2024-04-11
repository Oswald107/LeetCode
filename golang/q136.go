func singleNumber(nums []int) int {
	var m map[int]int
	m = make(map[int]int)
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}