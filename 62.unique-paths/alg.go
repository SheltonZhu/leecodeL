package _2_unique_paths

//func uniquePaths(m int, n int) int {
//	ways := 0
//	v := m - 1
//	h := n - 1
//	var backtrack func(v, h int)
//	backtrack = func(v, h int) {
//		if v == 0 || h == 0 {
//			ways++
//			return
//		}
//		backtrack(v-1, h)
//		backtrack(v, h-1)
//	}
//	backtrack(v, h)
//	return ways
//}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//func uniquePaths(m, n int) int {
//	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
//}
