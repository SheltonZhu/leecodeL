package removeduplicatesfromsortedarray2

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	left := 2
	for i := 2; i < n; i++ {
		if nums[left-2] != nums[i] {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}
