package _48_find_all_numbers_disappeared_in_an_array

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	output := []int{5, 6}
	ret := findDisappearedNumbers(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{1, 1}
	output := []int{2}
	ret := findDisappearedNumbers(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
