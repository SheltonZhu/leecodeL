package _28_longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	numsMap := map[int]bool{}
	// 去重
	for _, num := range nums {
		numsMap[num] = true
	}
	longestStreak := 0
	for num := range numsMap {
		if !numsMap[num-1] {
			currentStreak := 1
			for key := num + 1; numsMap[key]; key++ {
				currentStreak++
			}
			longestStreak = max(longestStreak, currentStreak)
		}
	}

	return longestStreak
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
