package _5_sort_colors

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	input := []int{2, 0, 2, 1, 1, 0}
	output := []int{0, 0, 1, 1, 2, 2}
	sortColors(input)
	if !reflect.DeepEqual(output, input) {
		t.Errorf("failed. output: %v, return: %v", output, input)
	}
}

func Test2(t *testing.T) {
	input := []int{2, 0, 1}
	output := []int{0, 1, 2}
	sortColors(input)
	if !reflect.DeepEqual(output, input) {
		t.Errorf("failed. output: %v, return: %v", output, input)
	}
}

func Test3(t *testing.T) {
	input := []int{0}
	output := []int{0}
	sortColors(input)
	if !reflect.DeepEqual(output, input) {
		t.Errorf("failed. output: %v, return: %v", output, input)
	}
}

func Test4(t *testing.T) {
	input := []int{1}
	output := []int{1}
	sortColors(input)
	if !reflect.DeepEqual(output, input) {
		t.Errorf("failed. output: %v, return: %v", output, input)
	}
}
