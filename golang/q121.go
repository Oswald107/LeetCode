func maxProfit(prices []int) int {
	max := 0
	cur := prices[0]
	for _, v := range prices {
		if v < cur {
			cur = v
		} else {
			if v-cur > max {
				max = v - cur
			}
		}
	}

	return max
}