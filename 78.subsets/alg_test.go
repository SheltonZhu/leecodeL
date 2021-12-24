package _8_subsets

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	input := []int{1, 2, 3}
	output := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
	ret := subsets(input)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	input := []int{0}
	output := [][]int{{}, {0}}
	ret := subsets(input)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
