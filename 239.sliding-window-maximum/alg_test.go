package _39_sliding_window_maximum

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	output := []int{3, 3, 5, 5, 6, 7}
	ret := maxSlidingWindow(nums, k)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{1}
	k := 1
	output := []int{1}
	ret := maxSlidingWindow(nums, k)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	nums := []int{1, -1}
	k := 1
	output := []int{1, -1}
	ret := maxSlidingWindow(nums, k)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test4(t *testing.T) {
	nums := []int{9, 11}
	k := 2
	output := []int{11}
	ret := maxSlidingWindow(nums, k)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test5(t *testing.T) {
	nums := []int{4, -2}
	k := 2
	output := []int{4}
	ret := maxSlidingWindow(nums, k)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
