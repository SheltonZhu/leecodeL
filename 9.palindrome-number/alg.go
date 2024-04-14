package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	ret := 0
	for x > ret {
		ret = ret*10 + x%10
		x = x / 10
	}
	return ret == x || x == ret/10
}
