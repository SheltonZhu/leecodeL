package _00_longestincreasing_subsequence

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	output := 4
	ret := lengthOfLIS(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test1(t *testing.T) {
	nums := []int{0, 1, 0, 3, 2, 3}
	output := 4
	ret := lengthOfLIS(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{7, 7, 7, 7, 7, 7, 7}
	output := 1
	ret := lengthOfLIS(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
