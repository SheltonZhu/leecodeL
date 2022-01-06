package _15_kth_largest_element_in_an_array

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	output := 5
	ret := findKthLargest(nums, k)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	output := 4
	ret := findKthLargest(nums, k)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
