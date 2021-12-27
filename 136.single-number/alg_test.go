package _36_single_number

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{2, 2, 1}
	output := 1
	ret := singleNumber(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	output := 4
	ret := singleNumber(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test3(t *testing.T) {
	nums := []int{1}
	output := 1
	ret := singleNumber(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
