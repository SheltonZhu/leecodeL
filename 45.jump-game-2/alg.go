package jumpgame2

func jump(nums []int) int {
	end := 0
	maxP := 0
	ans := 0
	for i, step := range nums[:len(nums)-1] {
		maxP = max(maxP, i+step)
		if i == end {
			end = maxP
			ans++
		}
	}
	return ans
}
