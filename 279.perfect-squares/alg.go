package _79_perfect_squares

import "math"

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		mmin := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			mmin = min(mmin, f[i-j*j])
		}
		f[i] = mmin + 1
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
