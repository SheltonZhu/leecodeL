package _98_house_robber

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	output := 4
	ret := rob(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	output := 12
	ret := rob(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
