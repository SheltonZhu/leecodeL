package _26_invert_binary_tree

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	root := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{1, nil, nil}, Right: &TreeNode{3, nil, nil},
		},
		Right: &TreeNode{7,
			&TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}},
	}
	output := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 7,
			Left: &TreeNode{9, nil, nil}, Right: &TreeNode{6, nil, nil},
		},
		Right: &TreeNode{2,
			&TreeNode{3, nil, nil}, &TreeNode{1, nil, nil}},
	}
	ret := invertTree(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
