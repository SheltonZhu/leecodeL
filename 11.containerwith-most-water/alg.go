package _1_containerwith_most_water

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxAreaValue := 0
	for left < right {
		area := Min(height[left], height[right]) * (right - left)
		maxAreaValue = Max(maxAreaValue, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxAreaValue
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
