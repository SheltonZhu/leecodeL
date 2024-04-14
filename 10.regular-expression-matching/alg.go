package regularexpressionmatching

// f[i][j] 表示 s 的前 i 个字符与 p 中的前 j 个字符是否能够匹配
// f[i][j] = f[i-1][j-1] && s[i] = p[j]

func isMatch(s string, p string) bool {
	// m, n := len(s), len(p)
	// matches := func(i, j int) bool {
	// 	if i == 0 {
	// 		return false
	// 	}
	// 	if p[j-1] == '.' {
	// 		return true
	// 	}
	// 	return s[i-1] == p[j-1]
	// }

	// f := make([][]bool, m+1)
	// for i := 0; i < len(f); i++ {
	// 	f[i] = make([]bool, n+1)
	// }
	// f[0][0] = true
	// for i := 0; i <= m; i++ {
	// 	for j := 1; j <= n; j++ {
	// 		if p[j-1] == '*' {
	// 			f[i][j] = f[i][j] || f[i][j-2]
	// 			if matches(i, j-1) {
	// 				f[i][j] = f[i][j] || f[i-1][j]
	// 			}
	// 		} else if matches(i, j) {
	// 			f[i][j] = f[i][j] || f[i-1][j-1]
	// 		}
	// 	}
	// }
	// return f[m][n]
	return false
}
