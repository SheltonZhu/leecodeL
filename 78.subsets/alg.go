package _8_subsets

func subsets(nums []int) [][]int {
	var result [][]int
	var backtrack func(subset []int)
	backtrack = func(subset []int) {
		result = append(result, append([]int{}, subset...))
		for i := 0; i < len(nums); i++ {
			if len(subset) == 0 || subset[len(subset)-1] < nums[i] {
				subset = append(subset, nums[i])
				backtrack(subset)
				subset = subset[:len(subset)-1]
			}
		}
	}
	backtrack([]int{})
	return result
}
