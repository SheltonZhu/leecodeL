package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left := 1
	for i := 1; i < n; i++ {
		if nums[left-1] != nums[i] {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}
