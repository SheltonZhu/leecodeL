package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[left] != nums[i] {
			nums[left+1] = nums[i]
			left++
		}
	}
	return left + 1
}
