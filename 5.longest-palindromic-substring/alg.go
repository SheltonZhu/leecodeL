package __longest_palindromic_substring

func longestPalindrome(s string) string {
	var (
		len, strLen      = 1, len(s)
		left, right      = 0, 0
		maxStart, maxLen = 0, 0
	)

	for i := 0; i < strLen; i++ {
		left = i - 1
		right = i + 1
		for left >= 0 && s[left] == s[i] {
			left--
			len++
		}
		for right < strLen && s[right] == s[i] {
			right++
			len++
		}
		for left >= 0 && right < strLen && s[left] == s[right] {
			left--
			right++
			len = len + 2
		}
		if len > maxLen {
			maxLen = len
			maxStart = left + 1
		}
		len = 1
	}
	return s[maxStart : maxStart+maxLen]
}

// i < j
// P(i,j) 表示字符串 s 的第 i 到 j 个字母组成的串
// P(i,j) = true , s[i,j] 是回文字符串
// P(i,j) = P(i+1,j-1) && s[i] == s[j]
// 长度为1或者2
// P(i,i) = true
// P(i,i+1) = s[i] == s[j]
// return s[i, max(j-i+1)]
// func longestPalindrome(s string) string {
// 	if s == "" || len(s) == 0 {
// 		return ""
// 	}
// 	strLen := len(s)
// 	dp := make([][]bool, strLen)
// 	for i := range dp {
// 		dp[i] = make([]bool, strLen)
// 	}
// 	maxLen := 1
// 	maxStart := 0
// 	for j := 1; j < strLen; j++ {
// 		for i := 0; i <= j; i++ {
// 			if s[i] == s[j] && (j-i <= 1 || (i < j && dp[i+1][j-1])) {
// 				dp[i][j] = true
// 				if j-i+1 > maxLen {
// 					maxLen = j - i + 1
// 					maxStart = i
// 				}
// 			}
// 		}
// 	}
// 	return s[maxStart : maxStart+maxLen]
// }
