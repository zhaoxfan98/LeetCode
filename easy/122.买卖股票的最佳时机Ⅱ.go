package easy

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			max = max + prices[i+1] - prices[i]
		}
	}
	return max
}
