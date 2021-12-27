package _04_maximum_depth_of_binary_tree

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	root := &TreeNode{
		3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil}},
	}
	output := 3
	ret := maxDepth(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
