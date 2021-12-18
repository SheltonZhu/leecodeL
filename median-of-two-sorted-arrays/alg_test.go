package median_of_two_sorted_arrays

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	ret := findMedianSortedArrays(nums1, nums2)
	expected := 2.0
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test2(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	ret := findMedianSortedArrays(nums1, nums2)
	expected := 2.5
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
func Test3(t *testing.T) {
	nums1 := []int{0, 0}
	nums2 := []int{0, 0}
	ret := findMedianSortedArrays(nums1, nums2)
	expected := 0.0
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test4(t *testing.T) {
	var nums1 []int
	nums2 := []int{1}
	ret := findMedianSortedArrays(nums1, nums2)
	expected := 1.0
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test5(t *testing.T) {
	nums1 := []int{2}
	var nums2 []int
	ret := findMedianSortedArrays(nums1, nums2)
	expected := 2.0
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
