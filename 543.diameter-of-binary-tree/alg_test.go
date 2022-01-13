package _43_diameter_of_binary_tree

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil},
		},
		&TreeNode{3, nil, nil},
	}
	output := 3
	ret := diameterOfBinaryTree(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{2, nil, nil},
		nil,
	}
	output := 1
	ret := diameterOfBinaryTree(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
