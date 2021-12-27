package _28_longest_consecutive_sequence

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	output := 4
	ret := longestConsecutive(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	output := 9
	ret := longestConsecutive(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
