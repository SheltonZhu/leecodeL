package mergesortedarray

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	output := []int{1, 2, 2, 3, 5, 6}
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(output, nums1) {
		t.Errorf("failed. output: %v, return: %v", output, nums1)
	}
}

func Test2(t *testing.T) {
	nums1 := []int{1}
	m := 1
	nums2 := []int{}
	n := 0

	output := []int{1}
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(output, nums1) {
		t.Errorf("failed. output: %v, return: %v", output, nums1)
	}
}

func Test3(t *testing.T) {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1

	output := []int{1}
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(output, nums1) {
		t.Errorf("failed. output: %v, return: %v", output, nums1)
	}
}

func Test4(t *testing.T) {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	m := 3
	nums2 := []int{1, 2, 3}
	n := 3

	output := []int{1, 2, 3, 4, 5, 6}
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(output, nums1) {
		t.Errorf("failed. output: %v, return: %v", output, nums1)
	}
}
