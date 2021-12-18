package two_sum

func twoSum(nums []int, target int) []int {
	posMap := map[int]int{}
	for idx, num := range nums {
		posMap[num] = idx
	}
	for idx, num := range nums {
		posArr := posMap[target-num]
		if posArr != 0 && posArr != idx {
			return []int{idx, posArr}
		}

	}
	return nil
}
