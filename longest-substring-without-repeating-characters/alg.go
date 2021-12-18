package longest_substring_without_repeating_characters

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	charMap := map[int32]int{}
	start := 0
	for idx, char := range s {
		if charMap[char] > 0 {
			start = int(math.Max(float64(charMap[char]), float64(start)))
		}
		maxLength = int(math.Max(float64(maxLength), float64(idx-start+1)))
		charMap[char] = idx + 1
	}
	return maxLength
}
