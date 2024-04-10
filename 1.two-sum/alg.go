package __two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for fi, num := range nums {
		if ti, ok := m[target-num]; ok {
			return []int{ti, fi}
		}
		m[num] = fi
	}
	return nil
}
