package _24_binary_tree_maximum_path_sum

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	output := 6
	ret := maxPathSum(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	root := &TreeNode{-10,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}
	output := 42
	ret := maxPathSum(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
