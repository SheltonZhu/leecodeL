package __longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	maxLen := 0
	start := 0
	for idx, char := range s {
		if v, ok := charMap[char]; ok && v > start {
			start = v
		}
		if idx-start+1 > maxLen {
			maxLen = idx - start + 1
		}
		charMap[char] = idx + 1
	}
	return maxLen
}
