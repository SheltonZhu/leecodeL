package _17_merge_two_binary_trees

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	root1 := &TreeNode{1,
		&TreeNode{3,
			&TreeNode{5, nil, nil},
			nil},
		&TreeNode{2, nil, nil},
	}
	root2 := &TreeNode{2,
		&TreeNode{1,
			nil,
			&TreeNode{4, nil, nil}},
		&TreeNode{3, nil,
			&TreeNode{7, nil, nil}},
	}
	output := &TreeNode{3,
		&TreeNode{4,
			&TreeNode{5, nil, nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{5, nil,
			&TreeNode{7, nil, nil}},
	}
	ret := mergeTrees(root1, root2)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test2(t *testing.T) {
	root1 := &TreeNode{1, nil, nil}
	root2 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	output := &TreeNode{2, &TreeNode{2, nil, nil}, nil}
	ret := mergeTrees(root1, root2)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
