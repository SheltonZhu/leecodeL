package _48_find_all_numbers_disappeared_in_an_array

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, num := range nums {
		num = (num - 1) % n
		nums[num] += n
	}
	var ret []int
	for i, num := range nums {
		if num <= n {
			ret = append(ret, i+1)
		}
	}
	return ret
}
