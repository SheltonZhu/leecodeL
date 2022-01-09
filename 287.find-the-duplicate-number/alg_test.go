package _87_find_the_duplicate_number

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}
	output := 2
	ret := findDuplicate(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{3, 1, 3, 4, 2}
	output := 3
	ret := findDuplicate(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	nums := []int{1, 1}
	output := 1
	ret := findDuplicate(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test4(t *testing.T) {
	nums := []int{1, 1, 2}
	output := 1
	ret := findDuplicate(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
