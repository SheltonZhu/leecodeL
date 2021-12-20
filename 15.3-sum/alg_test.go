package _5_3_sum

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	output := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	ret := threeSum(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	var nums []int
	var output [][]int
	ret := threeSum(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	nums := []int{0}
	var output [][]int
	ret := threeSum(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
