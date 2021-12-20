package __longest_palindromic_substring

//func longestPalindrome(s string) string {
//	if s == "" || len(s) == 0 {
//		return ""
//	}
//	strLen := len(s)
//	left := 0
//	right := 0
//	maxLen := 0
//	maxStart := 0
//	len := 1
//	for i := 0; i < strLen; i++ {
//		left = i - 1
//		right = i + 1
//		for left >= 0 && s[left] == s[i] {
//			left--
//			len++
//		}
//		for right < strLen && s[right] == s[i] {
//			right++
//			len++
//		}
//		for left >= 0 && right < strLen && s[left] == s[right] {
//			left--
//			right++
//			len = len + 2
//		}
//		if len > maxLen {
//			maxLen = len
//			maxStart = left
//		}
//		len = 1
//	}
//	return s[maxStart+1 : maxStart+1+maxLen]
//}
func longestPalindrome(s string) string {
	if s == "" || len(s) == 0 {
		return ""
	}
	strLen := len(s)
	dp := make([][]bool, strLen)
	for i := range dp {
		dp[i] = make([]bool, strLen)
	}
	maxEnd := 0
	maxStart := 0
	maxLen := 1
	for r := 1; r < strLen; r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-1 <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l+1 > maxLen {
					maxLen = r - l + 1
					maxEnd = r
					maxStart = l
				}
			}
		}
	}
	return s[maxStart : maxEnd+1]
}
