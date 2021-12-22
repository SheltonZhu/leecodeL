package _1_next_permutation

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	output := []int{1, 3, 2}
	if !reflect.DeepEqual(output, nums) {
		t.Errorf("failed. output: %v, return: %v", output, nums)
	}
}
func Test2(t *testing.T) {
	nums := []int{3, 2, 1}
	nextPermutation(nums)
	output := []int{1, 2, 3}
	if !reflect.DeepEqual(output, nums) {
		t.Errorf("failed. output: %v, return: %v", output, nums)
	}
}

func Test3(t *testing.T) {
	nums := []int{1, 1, 5}
	nextPermutation(nums)
	output := []int{1, 5, 1}
	if !reflect.DeepEqual(output, nums) {
		t.Errorf("failed. output: %v, return: %v", output, nums)
	}
}

func Test4(t *testing.T) {
	nums := []int{1}
	nextPermutation(nums)
	output := []int{1}
	if !reflect.DeepEqual(output, nums) {
		t.Errorf("failed. output: %v, return: %v", output, nums)
	}
}
