package two_sum

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	ret := twoSum(nums, target)
	expected := []int{0, 1}
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	ret := twoSum(nums, target)
	expected := []int{1, 2}
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	ret := twoSum(nums, target)
	expected := []int{0, 1}
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
func Test4(t *testing.T) {
	nums := []int{2, 5, 5, 11}
	target := 10
	ret := twoSum(nums, target)
	expected := []int{1, 2}
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test5(t *testing.T) {
	nums := []int{0, 4, 3, 0}
	target := 0
	ret := twoSum(nums, target)
	expected := []int{0, 3}
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test6(t *testing.T) {
	nums := []int{-1, -2, -3, -4, -5}
	target := -8
	ret := twoSum(nums, target)
	expected := []int{2, 4}
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
