package _05_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{3, nil, nil},
			&TreeNode{4, nil, nil},
		},
		&TreeNode{5, nil,
			&TreeNode{6, nil, nil}},
	}
	output := &TreeNode{1, nil,
		&TreeNode{2, nil,
			&TreeNode{3, nil,
				&TreeNode{4, nil,
					&TreeNode{5, nil,
						&TreeNode{6, nil, nil}}}},
		},
	}
	flatten(root)
	if !reflect.DeepEqual(output, root) {
		t.Errorf("failed. output: %v, return: %v", output, root)
	}
}

func Test2(t *testing.T) {
	root := &TreeNode{}
	output := &TreeNode{}
	flatten(root)
	if !reflect.DeepEqual(output, root) {
		t.Errorf("failed. output: %v, return: %v", output, root)
	}
}
