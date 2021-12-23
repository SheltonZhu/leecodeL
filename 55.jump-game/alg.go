package _5_jump_game

func canJump(nums []int) bool {
	last := len(nums) - 1
	if last == 0 {
		return true
	}
	far := 0
	for i := 0; i < last && i <= far; i++ {
		far = max(far, i+nums[i])
		if i+nums[i] >= last {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
