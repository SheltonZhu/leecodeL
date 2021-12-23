package _3_maximum_subarray

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	output := 6
	ret := maxSubArray(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test2(t *testing.T) {
	nums := []int{1}
	output := 1
	ret := maxSubArray(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test3(t *testing.T) {
	nums := []int{5, 4, -1, 7, 8}
	output := 23
	ret := maxSubArray(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
