func maxProfit(prices []int) int {
	output := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			output += prices[i] - prices[i-1]
		}
	}
	return output
}