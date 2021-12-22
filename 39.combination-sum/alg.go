package _9_combination_sum

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	//var backtrack func(target int, sum []int)
	//backtrack = func(target int, sum []int) {
	//	if target == 0 {
	//		result = append(result, append([]int{}, sum...))
	//		return
	//	}
	//	if target < 0 {
	//		return
	//	}
	//	for _, item := range candidates {
	//		if len(sum) == 0 || item >= sum[len(sum)-1] {
	//			sum = append(sum, item)
	//			backtrack(target-item, sum)
	//			sum = sum[:len(sum)-1]
	//		}
	//	}
	//}
	//
	//backtrack(target, []int{})
	backtrack(&result, candidates, target, []int{})
	return result
}

func backtrack(result *[][]int, candidates []int, target int, sum []int) {
	if target == 0 {
		*result = append(*result, append([]int{}, sum...))
		return
	}
	if target < 0 {
		return
	}
	for _, item := range candidates {
		if len(sum) == 0 || item >= sum[len(sum)-1] {
			sum = append(sum, item)
			backtrack(result, candidates, target-item, sum)
			sum = sum[:len(sum)-1]
		}
	}
}
