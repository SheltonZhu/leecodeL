package stringtointeger

import "math"

func myAtoi(s string) int {
	result, sign := 0, 1
	max := math.MaxInt32
	min := math.MinInt32
	i, n := 0, len(s)
	for ; i < n && s[i] == ' '; i++ {
	}
	if i >= n {
		return 0
	}

	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		sign = -1
	}

	for ; i < n; i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}

		result = result*10 + int(s[i]-'0')
		if sign*result < min {
			return min
		}
		if sign*result > max {
			return max
		}
	}

	return sign * result
}
