package _3_search_in_rotated_sorted_array

func searchRange(nums []int, target int) []int {
	leftmost := searchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := searchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}
func searchInts(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right + left) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

//func searchRange(nums []int, target int) []int {
//	if len(nums) == 1 {
//		if target == nums[0] {
//			return []int{0, 0}
//		}
//	}
//	left := 0
//	right := len(nums) - 1
//	for left <= right {
//		mid := (right + left) >> 1
//		if nums[mid] == target {
//			start, end := mid, mid
//			for start-1 >= 0 {
//				if nums[start-1] == target {
//					start = start - 1
//				} else {
//					break
//				}
//			}
//			for end+1 <= len(nums)-1 {
//				if nums[end+1] == target {
//					end = end + 1
//				} else {
//					break
//				}
//			}
//
//			return []int{start, end}
//		}
//		if nums[mid] > target && target >= nums[left] {
//			right = mid - 1
//		} else if nums[mid] < target && target <= nums[right] {
//			left = mid + 1
//		} else {
//			break
//		}
//	}
//	return []int{-1, -1}
//}
