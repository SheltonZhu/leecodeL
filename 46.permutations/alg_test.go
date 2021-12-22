package _6_permutations

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{1, 2, 3}
	output := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	ret := permute(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{0, 1}
	output := [][]int{{0, 1}, {1, 0}}
	ret := permute(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	nums := []int{1}
	output := [][]int{{1}}
	ret := permute(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
