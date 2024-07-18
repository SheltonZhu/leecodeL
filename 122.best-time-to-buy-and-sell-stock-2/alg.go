package besttimetobuyandsellstock2

func maxProfit(prices []int) int {
	pri := 0
	for i, p := range prices[1:] {
		if p-prices[i] > 0 {
			pri += p - prices[i]
		}
	}
	return pri
}
