package _87_find_the_duplicate_number

//func findDuplicate(nums []int) int {
//	slow, fast := 0, 0
//	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
//	}
//	slow = 0
//	for slow != fast {
//		slow = nums[slow]
//		fast = nums[fast]
//	}
//	return slow
//}
func findDuplicate(nums []int) int {
	numMap := map[int]bool{}
	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			return num
		}
		numMap[num] = true
	}
	return 0
}
