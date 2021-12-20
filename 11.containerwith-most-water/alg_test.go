package _1_containerwith_most_water

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	output := 49
	ret := maxArea(input)
	if reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	input := []int{1, 1}
	output := 1
	ret := maxArea(input)
	if reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	input := []int{4, 3, 2, 1, 4}
	output := 16
	ret := maxArea(input)
	if reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test4(t *testing.T) {
	input := []int{1, 2, 1}
	output := 2
	ret := maxArea(input)
	if reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
