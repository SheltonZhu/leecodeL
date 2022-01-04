package _98_house_robber

func rob(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	if numsLen == 1 {
		return nums[0]
	}
	//dp := make([]int, numsLen)
	//dp[0] = nums[0]
	//dp[1] = max(nums[1], nums[0])
	//for i := 2; i < numsLen; i++ {
	//	dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	//}
	//return dp[numsLen-1]
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
