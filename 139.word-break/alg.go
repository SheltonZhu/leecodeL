package _39_word_break

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	maxWordLen := 0
	for _, w := range wordDict {
		wordDictSet[w] = true
		maxWordLen = max(len(w), maxWordLen)
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if i-j > maxWordLen {
				continue
			}
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
