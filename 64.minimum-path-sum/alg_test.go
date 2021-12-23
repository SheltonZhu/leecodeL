package _4_minimum_path_sum

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	output := 7
	ret := minPathSum(grid)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	output := 12
	ret := minPathSum(grid)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
