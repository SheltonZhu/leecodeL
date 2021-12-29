package _8_rotate_image

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	output := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	rotate(matrix)
	if !reflect.DeepEqual(output, matrix) {
		t.Errorf("failed. output: %v, return: %v", output, matrix)
	}
}

func Test2(t *testing.T) {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	output := [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}
	rotate(matrix)
	if !reflect.DeepEqual(output, matrix) {
		t.Errorf("failed. output: %v, return: %v", output, matrix)
	}
}

func Test3(t *testing.T) {
	matrix := [][]int{{1}}
	output := [][]int{{1}}
	rotate(matrix)
	if !reflect.DeepEqual(output, matrix) {
		t.Errorf("failed. output: %v, return: %v", output, matrix)
	}
}

func Test4(t *testing.T) {
	matrix := [][]int{{1, 2}, {3, 4}}
	output := [][]int{{3, 1}, {4, 2}}
	rotate(matrix)
	if !reflect.DeepEqual(output, matrix) {
		t.Errorf("failed. output: %v, return: %v", output, matrix)
	}
}
