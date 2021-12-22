package _3_search_in_rotated_sorted_array

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	output := []int{3, 4}
	ret := searchRange(nums, target)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	output := []int{-1, -1}
	ret := searchRange(nums, target)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	var nums []int
	target := 0
	output := []int{-1, -1}
	ret := searchRange(nums, target)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test4(t *testing.T) {
	nums := []int{1}
	target := 1
	output := []int{0, 0}
	ret := searchRange(nums, target)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test5(t *testing.T) {
	nums := []int{2, 2}
	target := 1
	output := []int{-1, -1}
	ret := searchRange(nums, target)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test6(t *testing.T) {
	nums := []int{2, 2}
	target := 2
	output := []int{0, 1}
	ret := searchRange(nums, target)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test7(t *testing.T) {
	nums := []int{1, 4}
	target := 4
	output := []int{1, 1}
	ret := searchRange(nums, target)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
