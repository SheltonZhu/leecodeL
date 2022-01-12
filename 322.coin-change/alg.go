package _22_coin_change

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		mmin := math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 {
				mmin = min(mmin, dp[i-coin])
			}
		}
		dp[i] = mmin + 1
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
