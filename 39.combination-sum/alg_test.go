package _9_combination_sum

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	output := [][]int{{2, 2, 3}, {7}}
	ret := combinationSum(candidates, target)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8
	output := [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}
	ret := combinationSum(candidates, target)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test3(t *testing.T) {
	candidates := []int{2}
	target := 1
	var output [][]int
	ret := combinationSum(candidates, target)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test4(t *testing.T) {
	candidates := []int{1}
	target := 1
	output := [][]int{{1}}
	ret := combinationSum(candidates, target)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
