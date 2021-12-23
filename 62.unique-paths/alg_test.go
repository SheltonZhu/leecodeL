package _2_unique_paths

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	m, n := 3, 7
	output := 28
	ret := uniquePaths(m, n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	m, n := 3, 2
	output := 3
	ret := uniquePaths(m, n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test3(t *testing.T) {
	m, n := 7, 3
	output := 28
	ret := uniquePaths(m, n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test4(t *testing.T) {
	m, n := 3, 3
	output := 6
	ret := uniquePaths(m, n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test5(t *testing.T) {
	m, n := 51, 9
	output := 1916797311
	ret := uniquePaths(m, n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
