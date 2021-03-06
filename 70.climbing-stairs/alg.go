package _0_climbing_stairs

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//func climbStairs(n int) int {
//	p, q, r := 0, 0, 1
//	for i := 1; i <= n; i++ {
//		p = q
//		q = r
//		r = p + q
//	}
//	return r
//}
