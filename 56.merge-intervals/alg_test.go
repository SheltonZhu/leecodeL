package _6_merge_intervals

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	output := [][]int{{1, 6}, {8, 10}, {15, 18}}
	ret := merge(intervals)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	intervals := [][]int{{1, 4}, {4, 5}}
	output := [][]int{{1, 5}}
	ret := merge(intervals)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	intervals := [][]int{{1, 4}, {0, 4}}
	output := [][]int{{0, 4}}
	ret := merge(intervals)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
