package _6_permutations

func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(pool []int, combination []int)
	backtrack = func(pool []int, combination []int) {
		if len(nums) == len(combination) {
			result = append(result, append([]int{}, combination...))
			return
		}
		for i := 0; i < len(pool); i++ {
			combination = append(combination, pool[i])
			backtrack(append(append([]int{}, pool[:i]...), pool[i+1:]...), combination)
			combination = combination[:len(combination)-1]
		}
	}
	backtrack(nums, []int{})
	return result
}
