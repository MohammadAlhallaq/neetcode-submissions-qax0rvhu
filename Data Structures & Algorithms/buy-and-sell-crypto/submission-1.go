func maxProfit(prices []int) int {
	maxProfit := 0
	left := 0

	for right := 1; right < len(prices); right++ {

		if prices[right] > prices[left] {
			profit := prices[right] - prices[left]
			if profit > maxProfit {
				maxProfit = profit
			}
		}else{
			left = right
		}
	}
	return maxProfit
}
