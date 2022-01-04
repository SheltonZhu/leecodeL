package _69_majority_element

func majorityElement(nums []int) int {
	count := 0
	var candidate int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}

	}
	return candidate
}

//func majorityElement(nums []int) int {
//	maxCount := 0
//	maxNum := 0
//	numsMap := map[int]int{}
//	for _, num := range nums {
//		numsMap[num]++
//		if numsMap[num] > maxCount {
//			maxCount = numsMap[num]
//			maxNum = num
//		}
//	}
//	return maxNum
//}
