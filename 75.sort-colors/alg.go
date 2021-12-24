package _5_sort_colors

func sortColors(nums []int) {
	swapColor(nums[swapColor(nums, 0):], 1)
}

func swapColor(colors []int, target int) int {
	targets := 0
	for i, color := range colors {
		if color == target {
			colors[i], colors[targets] = colors[targets], colors[i]
			targets++
		}
	}
	return targets
}
