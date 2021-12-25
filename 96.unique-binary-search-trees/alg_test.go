package _6_unique_binary_search_trees

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	n := 3
	output := 5
	ret := numTrees(n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	n := 1
	output := 1
	ret := numTrees(n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
