package _21_best_time_to_buy_and_sell_stock

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfitValue := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			maxProfitValue = max(prices[i]-minPrice, maxProfitValue)
		}
	}
	return maxProfitValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
