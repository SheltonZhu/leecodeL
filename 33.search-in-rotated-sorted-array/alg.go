package _3_search_in_rotated_sorted_array

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) >> 1
		if target == nums[mid] {
			return mid
		}

		if nums[left] <= nums[mid] { // 左半有序，在左半查
			if target >= nums[left] && target < nums[mid] { // 在左半，二分左半
				right = mid - 1
			} else { // 不在左半，二分右半
				left = mid + 1
			}
		} else { // 右半有序，在右半查
			if target > nums[mid] && target <= nums[right] { // 在右半，二分右半
				left = mid + 1
			} else { // 不在右半，二分左半
				right = mid - 1
			}
		}
	}
	return -1
}
